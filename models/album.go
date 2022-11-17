package models

type Album struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Artist Artist `json:"artist"`
	Tracks Tracks `json:"tracks"`
}