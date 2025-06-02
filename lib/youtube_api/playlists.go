package youtubeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PlaylistItem struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ResourceID  struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"resourceId"`
	} `json:"snippet"`
}

type PlaylistItemResponse struct {
	Kind          string         `json:"kind"`
	Etag          string         `json:"etag"`
	NextPageToken string         `json:"nextPageToken,omitempty"`
	Items         []PlaylistItem `json:"items"`
}

type Playlist struct {
	apiKey string
}

func (p *Playlist) getOnePage(playlistId string) (*PlaylistItemResponse, error) {
	nextPageToken := ""
	var res *PlaylistItemResponse
	iter := 0

	for nextPageToken != "" || iter == 0 {
		iter++
		fmt.Printf("nextpage: %s, iter: %d\n", nextPageToken, iter)
		url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlistItems?key=%s&part=snippet&playlistId=%s&maxResults=50&pageToken=%s", p.apiKey, playlistId, nextPageToken)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response PlaylistItemResponse
		if err := json.Unmarshal(body, &response); err != nil {
			return nil, err
		}

		if res == nil {
			res = &response
		} else {
			res.Items = append(res.Items, response.Items...)

		}

		// get nextPageToken if it exists
		nextPageToken = response.NextPageToken

		if nextPageToken == "" || len(response.Items) == 0 {
			return res, nil
		}
	}

	return res, fmt.Errorf("something went wrong")
}

func (p *Playlist) GetPlaylistItems(playlistId string) (*PlaylistItemResponse, error) {

	return p.getOnePage(playlistId)
}
