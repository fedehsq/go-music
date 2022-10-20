package song

import (
	"go-music/music/artist"
)

type Song struct {
	Id       int
	Title    string
	Url      string
	Preview  string
	Duration int
	Rank     int
	Artist   artist.Artist
}
