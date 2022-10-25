package main

import (
	"fmt"
	"go-music/config"
	"go-music/views"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", views.Index).Methods("GET")
	r.HandleFunc("/search", views.Search).Methods("GET")
	// serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	srv := &http.Server{
		Handler: r,
		Addr:    strings.ReplaceAll(config.Conf.Server, "http://", ""),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server running on " + config.Conf.Server)
	log.Fatal(srv.ListenAndServe())
}
