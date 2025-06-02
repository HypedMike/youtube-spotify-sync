package spotifyapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Search struct {
	getToken func() *TokenResponse
}

func (s *Search) Search(query string) (*SearchResponse, error) {
	url := "https://api.spotify.com/v1/search"

	// substitute all spaces in the query with %20
	queryParams := strings.ReplaceAll(query, " ", "%2520")

	formattedUrl := fmt.Sprintf("%s?q=%s&type=track", url, queryParams)

	request, err := http.NewRequest("GET", formattedUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.getToken().AccessToken))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse *SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse, nil
}
