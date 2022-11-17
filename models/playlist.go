package models

type Playlist struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	PictureBig  string `json:"picture_big"`
	Duration    int    `json:"duration"`
	Nbtracks    int    `json:"nb_tracks"`
	Descritpion string `json:"description"`
	Tracks      Tracks `json:"tracks"`
}
