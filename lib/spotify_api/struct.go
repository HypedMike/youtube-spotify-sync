package spotifyapi

// Common types used across multiple structs
type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}

type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Artists              []Artist     `json:"artists"`
}

type Track struct {
	Album            Album    `json:"album"`
	Artists          []Artist `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber       int      `json:"disc_number"`
	DurationMs       int      `json:"duration_ms"`
	Explicit         bool     `json:"explicit"`
	ExternalIDs      struct {
		ISRC string `json:"isrc"`
		EAN  string `json:"ean"`
		UPC  string `json:"upc"`
	} `json:"external_ids"`
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	IsPlayable   bool         `json:"is_playable"`
	LinkedFrom   struct{}     `json:"linked_from"`
	Restrictions Restrictions `json:"restrictions"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	PreviewURL   string       `json:"preview_url"`
	TrackNumber  int          `json:"track_number"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	IsLocal      bool         `json:"is_local"`
}

type Playlist struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalURLs  ExternalURLs `json:"external_urls"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []Image      `json:"images"`
	Name          string       `json:"name"`
	Owner         struct {
		ExternalURLs ExternalURLs `json:"external_urls"`
		Href         string       `json:"href"`
		ID           string       `json:"id"`
		Type         string       `json:"type"`
		URI          string       `json:"uri"`
		DisplayName  string       `json:"display_name"`
	} `json:"owner"`
	Public     bool   `json:"public"`
	SnapshotID string `json:"snapshot_id"`
	Tracks     struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type PlaylistItemResponse struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalURLs  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name  string `json:"name"`
	Owner struct {
		ExternalURLs struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	Public     bool   `json:"public"`
	SnapshotID string `json:"snapshot_id"`
	Tracks     struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			AddedAt string `json:"added_at"`
			AddedBy struct {
				ExternalURLs struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"added_by"`
			IsLocal bool `json:"is_local"`
			Track   struct {
				Album struct {
					AlbumType        string   `json:"album_type"`
					TotalTracks      int      `json:"total_tracks"`
					AvailableMarkets []string `json:"available_markets"`
					ExternalURLs     struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href   string `json:"href"`
					ID     string `json:"id"`
					Images []struct {
						URL    string `json:"url"`
						Height int    `json:"height"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name                 string `json:"name"`
					ReleaseDate          string `json:"release_date"`
					ReleaseDatePrecision string `json:"release_date_precision"`
					Restrictions         struct {
						Reason string `json:"reason"`
					} `json:"restrictions"`
					Type    string `json:"type"`
					URI     string `json:"uri"`
					Artists []struct {
						ExternalURLs struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						ID   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						URI  string `json:"uri"`
					} `json:"artists"`
				} `json:"album"`
				Artists []struct {
					ExternalURLs struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				DiscNumber       int      `json:"disc_number"`
				DurationMS       int      `json:"duration_ms"`
				Explicit         bool     `json:"explicit"`
				ExternalIDs      struct {
					ISRC string `json:"isrc"`
					EAN  string `json:"ean"`
					UPC  string `json:"upc"`
				} `json:"external_ids"`
				ExternalURLs struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href       string `json:"href"`
				ID         string `json:"id"`
				IsPlayable bool   `json:"is_playable"`
				LinkedFrom struct {
				} `json:"linked_from"`
				Restrictions struct {
					Reason string `json:"reason"`
				} `json:"restrictions"`
				Name        string `json:"name"`
				Popularity  int    `json:"popularity"`
				PreviewURL  string `json:"preview_url"`
				TrackNumber int    `json:"track_number"`
				Type        string `json:"type"`
				URI         string `json:"uri"`
				IsLocal     bool   `json:"is_local"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type PlaylistList struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		Collaborative bool   `json:"collaborative"`
		Description   string `json:"description"`
		ExternalURLs  struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		ID     string `json:"id"`
		Images []struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name  string `json:"name"`
		Owner struct {
			ExternalURLs struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
			DisplayName string `json:"display_name"`
		} `json:"owner"`
		Public     bool   `json:"public"`
		SnapshotID string `json:"snapshot_id"`
		Tracks     struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"tracks"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"items"`
}

type Show struct {
	AvailableMarkets []string `json:"available_markets"`
	Copyrights       []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	Description        string       `json:"description"`
	HTMLDescription    string       `json:"html_description"`
	Explicit           bool         `json:"explicit"`
	ExternalURLs       ExternalURLs `json:"external_urls"`
	Href               string       `json:"href"`
	ID                 string       `json:"id"`
	Images             []Image      `json:"images"`
	IsExternallyHosted bool         `json:"is_externally_hosted"`
	Languages          []string     `json:"languages"`
	MediaType          string       `json:"media_type"`
	Name               string       `json:"name"`
	Publisher          string       `json:"publisher"`
	Type               string       `json:"type"`
	URI                string       `json:"uri"`
	TotalEpisodes      int          `json:"total_episodes"`
}

type Episode struct {
	AudioPreviewURL      string       `json:"audio_preview_url"`
	Description          string       `json:"description"`
	HTMLDescription      string       `json:"html_description"`
	DurationMs           int          `json:"duration_ms"`
	Explicit             bool         `json:"explicit"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	IsExternallyHosted   bool         `json:"is_externally_hosted"`
	IsPlayable           bool         `json:"is_playable"`
	Language             string       `json:"language"`
	Languages            []string     `json:"languages"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	ResumePoint          struct {
		FullyPlayed      bool `json:"fully_played"`
		ResumePositionMs int  `json:"resume_position_ms"`
	} `json:"resume_point"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	Restrictions Restrictions `json:"restrictions"`
}

type Audiobook struct {
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
	AvailableMarkets []string `json:"available_markets"`
	Copyrights       []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	Description     string       `json:"description"`
	HTMLDescription string       `json:"html_description"`
	Edition         string       `json:"edition"`
	Explicit        bool         `json:"explicit"`
	ExternalURLs    ExternalURLs `json:"external_urls"`
	Href            string       `json:"href"`
	ID              string       `json:"id"`
	Images          []Image      `json:"images"`
	Languages       []string     `json:"languages"`
	MediaType       string       `json:"media_type"`
	Name            string       `json:"name"`
	Narrators       []struct {
		Name string `json:"name"`
	} `json:"narrators"`
	Publisher     string `json:"publisher"`
	Type          string `json:"type"`
	URI           string `json:"uri"`
	TotalChapters int    `json:"total_chapters"`
}

// Paging is a generic struct for paginated results
type Paging struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

// SearchResponse represents the full search response from Spotify
type SearchResponse struct {
	Tracks struct {
		Paging
		Items []Track `json:"items"`
	} `json:"tracks"`
	Artists struct {
		Paging
		Items []Artist `json:"items"`
	} `json:"artists"`
	Albums struct {
		Paging
		Items []Album `json:"items"`
	} `json:"albums"`
	Playlists struct {
		Paging
		Items []Playlist `json:"items"`
	} `json:"playlists"`
	Shows struct {
		Paging
		Items []Show `json:"items"`
	} `json:"shows"`
	Episodes struct {
		Paging
		Items []Episode `json:"items"`
	} `json:"episodes"`
	Audiobooks struct {
		Paging
		Items []Audiobook `json:"items"`
	} `json:"audiobooks"`
}
