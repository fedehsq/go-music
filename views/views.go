package views

import (
	"encoding/json"
	"fmt"
	"go-music/music"
	"go-music/config"
	"html/template"
	"log"
	"net/http"
)

type Songs struct {
	Songs []music.Song `json:"data"`
}


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

func Index(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(config.Conf.TopItaliansURL)
	if err != nil {
		log.Fatalln(err)
	}
	t, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, createHtmlResponse(resp, "Top 50 Italia"))
}
