package api

import (
	"encoding/json"
	"fmt"
	"github.com/fedehsq/go-music/music"
	"github.com/fedehsq/go-music/config"
	"log"
	"net/http"
)

type Songs struct {
	Songs []music.Song `json:"data"`
}

func (s *Songs) getDownloadUrl() {
	for i := 0; i < len(s.Songs); i++ {
		s.Songs[i].Url = config.Conf.SongUrl + fmt.Sprintf("%d", s.Songs[i].Id)
	}
}

/*
func createHtmlResponse(r *http.Response, pageTitle string) map[string]interface{} {
	var songs Songs
	json.NewDecoder(r.Body).Decode(&songs)
	for i := 0; i < len(songs.Songs); i++ {
		url := config.Conf.SongUrl + fmt.Sprintf("%d", songs.Songs[i].Id)
		songs.Songs[i].Url = url
	}
	return map[string]interface{}{"songs": songs.Songs, "title": pageTitle}
}

func Search(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("search")
	resp, err := http.Get(config.Conf.SearchURL + data)
	if err != nil {
		log.Fatalln(err)
	}
	t, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, createHtmlResponse(resp,"Risultati ricerca per: " + data))
}
*/


func GetTopItalians(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(config.Conf.TopItaliansURL)
	if err != nil {
		log.Fatalln(err)
	}
	var songs Songs
	json.NewDecoder(resp.Body).Decode(&songs)
	// Apply the url to each song without looping
	songs.getDownloadUrl()
	// write the response as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}
