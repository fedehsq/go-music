package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fedehsq/go-music/api"
	"github.com/fedehsq/go-music/config"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	r.HandleFunc("/playlist-songs", api.GetPlaylistSongs).Methods("GET")
	handler := c.Handler(r)
	//r.HandleFunc("/search", views.Search).Methods("GET")
	// serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	srv := &http.Server{
		Handler:      handler,
		Addr:         strings.ReplaceAll(config.Conf.Server, "http://", ""),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server running on " + config.Conf.Server)
	log.Fatal(srv.ListenAndServe())
}
