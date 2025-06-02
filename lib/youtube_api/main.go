package youtubeapi

type YouTubeAPI struct {
	APIKey   string
	Playlist *Playlist
}

func NewYouTubeAPI(apiKey string) *YouTubeAPI {
	playList := &Playlist{apiKey: apiKey}

	return &YouTubeAPI{APIKey: apiKey, Playlist: playList}
}
