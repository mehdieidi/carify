package divar

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	kithttp "back/pkg/kittransport/http"
)

var (
	divarGetURL, _ = url.Parse("https://api.divar.ir/v8/posts-v2/web")
)

type GetRequest struct {
	CarToken string
}

func validateGetRequest(req any) error {
	request, ok := req.(GetRequest)
	if !ok {
		return errors.New("invalid get request")
	}
	if request.CarToken == "" {
		return errors.New("invalid car token")
	}

	return nil
}

type Widget struct {
	WidgetType string `json:"widget_type"`
	Data       struct {
		Items []struct {
			Title string `json:"title"`
			Value string `json:"value"`
		} `json:"items,omitempty"`
		Title string `json:"title,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"data"`
}

type Section struct {
	SectionName string   `json:"section_name"`
	Widgets     []Widget `json:"widgets"`
}

type GetResponse struct {
	Sections []Section `json:"sections"`
}

func createGetRequest(ctx context.Context, req any) (*http.Request, error) {
	if err := validateGetRequest(req); err != nil {
		return nil, err
	}

	getRequest := req.(GetRequest)

	link := divarGetURL.JoinPath(getRequest.CarToken)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, link.String(), nil)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func decodeGetResponse(ctx context.Context, res *http.Response) (any, error) {
	defer res.Body.Close()

	var resp GetResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func Get(ctx context.Context, carToken string) (GetResponse, error) {
	httpClient := http.DefaultClient
	httpClient.Timeout = 10 * time.Second

	client := kithttp.NewExplicitClient(
		createGetRequest,
		decodeGetResponse,
		kithttp.SetClient(httpClient),
		kithttp.ClientBefore(
			kithttp.SetRequestHeader("Content-Type", "application/json"),
		),
	)

	res, err := client.Endpoint()(ctx, GetRequest{CarToken: carToken})
	if err != nil {
		return GetResponse{}, err
	}

	resp, ok := res.(GetResponse)
	if !ok {
		return GetResponse{}, errors.New("cant assert get response")
	}

	return resp, nil
}
