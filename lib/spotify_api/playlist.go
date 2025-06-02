package spotifyapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type PlaylistActions struct {
	ID       string
	getToken func() *TokenResponse
}

type Options struct {
	MaxLimit int
}

func (p *PlaylistActions) GetPlaylistItems(playlistID string, options Options) (*PlaylistItemResponse, error) {

	if options.MaxLimit == 0 {
		options.MaxLimit = 50
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s?maxResults=%d", playlistID, options.MaxLimit)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.getToken().AccessToken))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var playlistItemResponse *PlaylistItemResponse
	err = json.Unmarshal(body, &playlistItemResponse)
	if err != nil {
		return nil, err
	}

	return playlistItemResponse, nil
}

func (p *PlaylistActions) GetUsersPlaylists(userID string) (*PlaylistList, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", userID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.getToken().AccessToken))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var playlistItemResponse *PlaylistList
	err = json.Unmarshal(body, &playlistItemResponse)
	if err != nil {
		return nil, err
	}

	return playlistItemResponse, nil
}

func (p *PlaylistActions) RemovePlaylistItem(playlistID string, itemIDs []string) error {

	// build the body
	type bodyStruct struct {
		Tracks []struct {
			URI string `json:"uri"`
		} `json:"tracks"`
	}

	var tracks bodyStruct
	for _, itemID := range itemIDs {
		tracks.Tracks = append(tracks.Tracks, struct {
			URI string `json:"uri"`
		}{
			URI: itemID,
		})
	}

	body, err := json.Marshal(tracks)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistID)

	request, err := http.NewRequest("DELETE", url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.getToken().AccessToken))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		// unmarshal the body
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed to remove playlist item: %s, %s, %s", response.Status, string(body), string(responseBody))
	}

	return nil
}

func (p *PlaylistActions) AddPlaylistItem(playlistID string, itemID string) error {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistID)

	var bodyStruct struct {
		URIs []string `json:"uris"`
	}
	bodyStruct.URIs = []string{itemID}
	body, err := json.Marshal(bodyStruct)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.getToken().AccessToken))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 201 {
		return fmt.Errorf("failed to add playlist item: %s", response.Status)
	}

	return nil
}
