package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"go-music/views"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", views.Index).Methods("GET")
	r.HandleFunc("/search", views.Search).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "192.168.1.4:19090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started at port 19090")
	fmt.Println("http://192.168.1.4:19090")
	log.Fatal(srv.ListenAndServe())
}
