package models

type ChartPlaylist struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	PictureBig string `json:"picture_big"`
	Link       string `json:"link"`
	Tracklist  string `json:"tracklist"`
}

type ChartPlaylists struct {
	Playlists []ChartPlaylist `json:"data"`
}
