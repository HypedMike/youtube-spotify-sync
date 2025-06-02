# Youtube Spotify sync

This is a utility that i created to help me out with an everyday task. Basically my car only read Spotify but I use
Youtube Music for most of my stuff.

I created this utility to keep two playlist synced. The process is simple:
1 - deletes all the songs from a specific spotify playlist
2 - get all the songs from a specific Youtube Playlist
3 - search all the songs from YouTube on Spotify
4 - Add all the results in the Spotify Playlist

# Requiriments
- Google Cloud Api Key with Data API enabled
- Spotify API Key
- a fallback website for Spotify Authentication

# Set up
This is what the env file should look like

```bash
YOUTUBE_API_KEY=<secret>
SPOTIFY_API_KEY=<secret>
SPOTIFY_CLIENT_ID=<secret>
YOUTUBE_PLAYLIST_ID=<secret>
SPOTIFY_FALLBACK=<secret>
```