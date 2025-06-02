package main

import (
	"fmt"
	"os"
	spotifyapi "yt-spotify-sync/lib/spotify_api"
	youtubeapi "yt-spotify-sync/lib/youtube_api"

	"github.com/joho/godotenv"
)

const SPOTIFY_PLAYLIST_ID = "1sFkRepDrO7WrhmfYsUFhS"

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	spotifyAPI := spotifyapi.NewSpotifyAPI(os.Getenv("SPOTIFY_API_KEY"), os.Getenv("SPOTIFY_CLIENT_ID"))

	spotifyAPI.Auth.Authorize()

	spotifyPlaylistItems, err := spotifyAPI.Playlist.GetPlaylistItems(SPOTIFY_PLAYLIST_ID, spotifyapi.Options{MaxLimit: 50})
	if err != nil {
		fmt.Printf("Error getting playlist items: %v\n", err)
		return
	}

	// remove all items from the spotify playlist
	itemIDs := []string{}
	for _, item := range spotifyPlaylistItems.Tracks.Items {
		itemIDs = append(itemIDs, item.Track.URI)
	}
	if len(itemIDs) > 0 {
		err = spotifyAPI.Playlist.RemovePlaylistItem(SPOTIFY_PLAYLIST_ID, itemIDs)
		if err != nil {
			fmt.Printf("Error removing playlist items: %v\n", err)
			return
		}
	}

	fmt.Println("Playlist items removed successfully")

	youtubeAPI := youtubeapi.NewYouTubeAPI(os.Getenv("YOUTUBE_API_KEY"))

	youtubePlaylistItems, err := youtubeAPI.Playlist.GetPlaylistItems(os.Getenv("YOUTUBE_PLAYLIST_ID"))
	if err != nil {
		fmt.Printf("Error getting playlist items: %v\n", err)
		return
	}

	for _, item := range youtubePlaylistItems.Items {

		fmt.Printf("Adding item %s\n", item.Snippet.Title)

		// search for the item in spotify
		spotifyItem, err := spotifyAPI.Search.Search(item.Snippet.Title)
		if err != nil {
			fmt.Printf("Error searching for item %s: %v\n", item.Snippet.Title, err)
			continue
		}

		if len(spotifyItem.Tracks.Items) == 0 {
			fmt.Printf("No spotify item found for %s\n", item.Snippet.Title)
			continue
		}

		err = spotifyAPI.Playlist.AddPlaylistItem(SPOTIFY_PLAYLIST_ID, spotifyItem.Tracks.Items[0].URI)
		if err != nil {
			fmt.Printf("Error adding playlist item %s: %v\n", item.Snippet.Title, err)
			continue
		}
	}

	fmt.Println("All playlist items added successfully")
}
