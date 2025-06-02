package spotifyapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Auth struct {
	apiKey      string
	clientID    string
	accessToken *TokenResponse
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (a *Auth) generateRandomString(length int) string {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = characters[r.Intn(len(characters))]
	}
	return string(b)
}

func (a *Auth) Authorize() error {
	url := "https://accounts.spotify.com/authorize?response_type=code&client_id=" + a.clientID +
		"&scope=playlist-modify-public%20playlist-modify-private&redirect_uri=https://classicum.fly.dev&state=" +
		a.generateRandomString(10) + "&show_dialog=false"

	fmt.Println("Open this URL in your browser: " + url)

	// wait for user input
	fmt.Println("After authorization, please paste the code from the URL here:")
	var code string
	fmt.Scanln(&code)

	// Get the access token using the authorization code
	tokenResponse, err := a.getTokenWithCode(code)
	if err != nil {
		return err
	}

	a.accessToken = tokenResponse

	return nil
}

func (a *Auth) getToken() (*TokenResponse, error) {
	url := "https://accounts.spotify.com/api/token"

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	bodyReq := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", a.clientID, a.apiKey)

	request.Body = io.NopCloser(strings.NewReader(bodyReq))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse *TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	a.accessToken = tokenResponse

	return tokenResponse, nil
}

func (a *Auth) getTokenWithCode(code string) (*TokenResponse, error) {
	url := "https://accounts.spotify.com/api/token"

	bodyReq := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=https://classicum.fly.dev&client_id=%s&client_secret=%s",
		code, a.clientID, a.apiKey)

	request, err := http.NewRequest("POST", url, strings.NewReader(bodyReq))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse *TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func (a *Auth) GetAccessToken() *TokenResponse {
	if a.accessToken == nil {
		accessToken, err := a.getToken()
		if err != nil {
			return nil
		}

		a.accessToken = accessToken
	}

	return a.accessToken
}
