package spotifyapi

type SpotifyAPI struct {
	APIKey   string
	ClientID string
	Auth     *Auth
	Search   *Search
	Playlist *PlaylistActions
}

func NewSpotifyAPI(apiKey string, clientID string) *SpotifyAPI {
	auth := &Auth{apiKey: apiKey, clientID: clientID}
	search := &Search{getToken: auth.GetAccessToken}
	playlist := &PlaylistActions{getToken: auth.GetAccessToken}

	return &SpotifyAPI{APIKey: apiKey, ClientID: clientID, Auth: auth, Search: search, Playlist: playlist}
}
