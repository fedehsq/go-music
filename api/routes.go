package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fedehsq/go-music/models"
)

const (
	playlistUrl      = "https://api.deezer.com/playlist"
	tracksUrl        = "https://api.deezer.com/track"
	artistUrl        = "https://api.deezer.com/artist"
	albumUrl         = "https://api.deezer.com/album"
	chartPlaylistsUrl = "https://api.deezer.com/chart/0/playlists"
)

func Playlist(w http.ResponseWriter, r *http.Request) {
	// Get the playlist id from the url
	playlistId := r.URL.Query().Get("id")
	resp, err := http.Get(fmt.Sprintf("%s/%s", playlistUrl, playlistId))
	if err != nil {
		log.Fatalln(err)
	}
	var playlist models.Playlist
	json.NewDecoder(resp.Body).Decode(&playlist)
	playlist.Tracks.SetDownloadUrls()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlist)
}

func ChartPlaylists(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(chartPlaylistsUrl)
	if err != nil {
		log.Fatalln(err)
	}
	chartPlaylists := models.ChartPlaylists{}
	json.NewDecoder(resp.Body).Decode(&chartPlaylists)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chartPlaylists.Playlists)
}
