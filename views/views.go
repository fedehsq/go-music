package views

import (
	"encoding/json"
	//"fmt"
	"go-music/music/song"
	"html/template"
	"log"
	"net/http"
)

type Songs struct {
	Data []song.Song
}

type SongStruct struct {
	Force string
}

const (
	searchUrl   = "https://api.deezer.com/search?q="
	topItalians = "https://api.deezer.com/playlist/65489479/tracks"
)

func createHtmlResponse(r *http.Response, pageTitle string) map[string]interface{} {
	var songs Songs
	json.NewDecoder(r.Body).Decode(&songs)
	// for i := 0; i < len(songs.Data); i++ {
	// 	url := songUrl + fmt.Sprintf("%d", songs.Data[i].Id)
	// 	songs.Data[i].Url = url
	// }
	return map[string]interface{}{"songs": songs.Data, "title": pageTitle}
}

func Search(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("search")
	resp, err := http.Get(searchUrl + data)
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
	resp, err := http.Get(topItalians)
	if err != nil {
		log.Fatalln(err)
	}
	t, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, createHtmlResponse(resp, "Top 50 Italia"))
}
