package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "back/docs"
	"back/pkg/log"
	"back/services/predict"
	"back/services/site/setting"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	httpswagger "github.com/swaggo/http-swagger"
)

// @title    carify Backend API
// @BasePath /v1
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(os.Getenv("LOG_NAME"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile)

	settingFile, err := os.Open("./settings/settings.json")
	if err != nil {
		panic(err)
	}

	settingStr, err := io.ReadAll(settingFile)
	if err != nil {
		panic(err)
	}
	settingFile.Close()

	// Storage (repository) layer.

	// Service layer.
	settingService := setting.NewService(string(settingStr))
	predictService := predict.NewService()

	// Transport layer.
	mux := chi.NewRouter()

	mux.Use(cors.AllowAll().Handler)

	mux.Get("/docs/swagger/*", httpswagger.Handler())

	mux.Mount("/v1/site/settings", setting.MakeHTTPHandler(
		setting.LoggingMiddleware(logger)(settingService),
		logger,
	))
	mux.Mount("/v1/costs/predict", predict.MakeHTTPHandler(
		predict.LoggingMiddleware(logger)(predictService),
		logger,
	))

	srv := &http.Server{
		Addr:         os.Getenv("SERVER_ADDR"),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	errC := make(chan error, 1)

	go func() {
		<-ctx.Done()

		fmt.Println("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		fmt.Println("Shutdown completed")
	}()

	go func() {
		fmt.Printf("Listening and serving [%s]\n", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	if err := <-errC; err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("exiting...")
}
