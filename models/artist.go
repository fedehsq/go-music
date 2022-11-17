package models

type Artist struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	Tracklist string `json:"tracklist"`
}
