package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/schollz/progressbar/v3"

	"github.com/mehdieidi/carify/data/internal/database"
	"github.com/mehdieidi/carify/data/pkg/log"
	"github.com/mehdieidi/carify/data/protocol/derror"
	"github.com/mehdieidi/carify/data/services/car"
	"github.com/mehdieidi/carify/data/services/pcar"
	"github.com/mehdieidi/carify/data/services/preprocess"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := database.Open(
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PASSWORD"),
		),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := database.AutoMigrate(db); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(os.Getenv("LOG_NAME"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile)

	ctx := context.Background()

	// Storage (repository) layer.
	carStorage := car.NewStorage(db, logger)
	preProcessStorage := preprocess.NewStorage(db, logger)
	pCarStorage := pcar.NewStorage(db, logger)
	carService := car.NewService(carStorage, logger)
	preProcessService := preprocess.NewService(preProcessStorage, carService)
	pCarService := pcar.NewService(pCarStorage, preProcessService)

	fetchFlag := flag.Bool("fetch", false, "fetch data?")
	preProcessFlag := flag.Bool("preprocess", false, "preprocess?")
	oneHotFlag := flag.Bool("onehot", false, "one hot encode?")
	flag.Parse()

	if *fetchFlag {
		cities := strings.Split(os.Getenv("CITIES"), " ")
		brandModel := os.Getenv("BRAND_MODEL")

		fmt.Println("-- searching divar for tokens --")

		carTokens, err := carService.Search(ctx, cities, brandModel)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%+v car tokens found]\n", len(carTokens))

		fmt.Printf("\n-- retrieving data --\n")
		pb := progressbar.Default(int64(len(carTokens)))

		for _, t := range carTokens {
			_, err := carStorage.FindByToken(ctx, t)
			if err == nil {
				pb.Add(1)
				continue
			}
			if err != nil {
				if !errors.Is(err, derror.ErrUnknownCar) {
					pb.Add(1)
					logger.Error("main", log.ServiceLayer, "main", log.Args{log.LogErrKey: err, "car_token": t})
					continue
				}
			}

			time.Sleep(1 * time.Second)

			c, err := carService.Get(ctx, t)
			if err != nil {
				pb.Add(1)
				logger.Error("main", log.ServiceLayer, "main", log.Args{log.LogErrKey: err, "car_token": t})
				continue
			}

			_, err = carStorage.Store(ctx, c)
			if err != nil {
				pb.Add(1)
				logger.Error("main", log.ServiceLayer, "main", log.Args{log.LogErrKey: err, "car_token": t})
				continue
			}

			pb.Add(1)
		}
	}

	if *preProcessFlag {
		err := preProcessService.Year(ctx, 1389, 1400)
		if err != nil {
			panic(err)
		}

		err = preProcessService.Color(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.UsageKM(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.BodyStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.CashCost(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.MotorStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.RearChassisStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.FrontChassisStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.InsuranceDue(ctx)
		if err != nil {
			panic(err)
		}

		err = preProcessService.GearBox(ctx)
		if err != nil {
			panic(err)
		}
	}

	if *oneHotFlag {
		pCarService.OneHotEncode(ctx)
	}
}
