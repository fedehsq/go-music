package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fedehsq/go-music/config"
	"github.com/fedehsq/go-music/music"
)

type Songs struct {
	Songs []music.Song `json:"data"`
}

func (s *Songs) getDownloadUrl() {
	for i := 0; i < len(s.Songs); i++ {
		s.Songs[i].Url = config.Conf.SongUrl + fmt.Sprintf("%d", s.Songs[i].Id)
	}
}

func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	// Get the playlist id from the url
	playlistId := r.URL.Query().Get("id")
	resp, err := http.Get(fmt.Sprintf("%s/%s/tracks", config.Conf.PlaylistUrl, playlistId))
	if err != nil {
		log.Fatalln(err)
	}
	var songs Songs
	json.NewDecoder(resp.Body).Decode(&songs)
	songs.getDownloadUrl()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs.Songs)
}
