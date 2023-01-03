package divar

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	kithttp "bitbucket.imenaria.org/tool/toolkit/transport/http"
)

var (
	divarSearchURL, _ = url.Parse("https://api.divar.ir/v8/web-search/1/light")
)

type SearchRequest struct {
	JSONSchema struct {
		Category struct {
			Value string `json:"value"`
		} `json:"category"`
		BrandModel struct {
			Value []string `json:"value"`
		} `json:"brand_model"`
		Cities []string `json:"cities"`
	} `json:"json_schema"`
	LastPostDate int64 `json:"last-post-date"`
}

func NewSearchRequest(category string, brandModel string, cities []string, lastPostDate int64) SearchRequest {
	var s SearchRequest

	s.JSONSchema.Category.Value = category
	s.JSONSchema.BrandModel.Value = []string{brandModel}
	s.JSONSchema.Cities = cities
	s.LastPostDate = lastPostDate

	return s
}

func validateSearchRequest(req any) error {
	request, ok := req.(SearchRequest)
	if !ok {
		return errors.New("invalid car search request")
	}

	if request.JSONSchema.Category.Value == "" {
		return errors.New("invalid category value")
	}

	if request.JSONSchema.BrandModel.Value == nil || len(request.JSONSchema.BrandModel.Value) == 0 {
		return errors.New("invalid brand model value")
	}

	if request.JSONSchema.Cities == nil || len(request.JSONSchema.Cities) == 0 {
		return errors.New("invalid cities")
	}

	return nil
}

type SearchResponse struct {
	LastPostDate int64 `json:"last_post_date"`
	WebWidgets   struct {
		PostList []struct {
			Data struct {
				Action struct {
					Payload struct {
						Token string `json:"token"`
					} `json:"payload"`
				} `json:"action"`
			} `json:"data"`
		} `json:"post_list"`
	} `json:"web_widgets"`
}

func createSearchRequest(ctx context.Context, req any) (*http.Request, error) {
	if err := validateSearchRequest(req); err != nil {
		return nil, err
	}

	searchRequest := req.(SearchRequest)

	body, err := json.Marshal(searchRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, divarSearchURL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return request, nil
}

func decodeSearchResponse(ctx context.Context, res *http.Response) (any, error) {
	defer res.Body.Close()

	var resp SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func Search(ctx context.Context, searchRequest SearchRequest) (SearchResponse, error) {
	httpClient := http.DefaultClient
	httpClient.Timeout = 10 * time.Second

	client := kithttp.NewExplicitClient(
		createSearchRequest,
		decodeSearchResponse,
		kithttp.SetClient(httpClient),
		kithttp.ClientBefore(
			kithttp.SetRequestHeader("Content-Type", "application/json"),
		),
	)

	res, err := client.Endpoint()(ctx, searchRequest)
	if err != nil {
		return SearchResponse{}, err
	}

	resp, ok := res.(SearchResponse)
	if !ok {
		return SearchResponse{}, errors.New("cant assert search response")
	}

	return resp, nil
}
