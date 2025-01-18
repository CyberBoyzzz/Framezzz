package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CyberBoyzzz/Framezzz/internal/model"
	"net/http"
)

// FetchComicFromAPI fetches a comic from the xkcd API and maps it to the Comic struct
func FetchComicFromAPI(ctx context.Context, id int) (model.Comic, error) {
	var comic model.Comic

	apiURL := fmt.Sprintf("https://xkcd.com/%s/info.0.json", id)

	response, err := MakeAPICall(http.MethodGet, apiURL)
	if err != nil {
		return comic, fmt.Errorf("failed to fetch comic from API: %w", err)
	}

	var apiResponse model.ExternalComicAPIResponse
	if err := json.Unmarshal([]byte(response), &apiResponse); err != nil {
		return comic, fmt.Errorf("failed to parse API response: %w", err)
	}

	comic = model.Comic{
		ID:       apiResponse.Num,
		Title:    apiResponse.SafeTitle,
		CoverURL: apiResponse.Img,
		Likes:    0,
	}

	return comic, nil
}
