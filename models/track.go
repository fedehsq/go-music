package models

import "fmt"

const (
	songUrl = "https://hub.ilill.li/?id="
)

type Track struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Preview  string `json:"preview"`
	Mp3      string `json:"file"`
	Duration int    `json:"duration"`
	Rank     int    `json:"rank"`
	Url      string `json:"url"`
	Artist   Artist `json:"artist"`
}

func (s *Track) setDownloadUrl() {
	s.Url = fmt.Sprintf("%s/%d", songUrl, s.Id)
}

type Tracks struct {
	Tracks []Track `json:"data"`
}

func (tracks *Tracks) SetDownloadUrls() {
	for i := 0; i < len(tracks.Tracks); i++ {
		tracks.Tracks[i].setDownloadUrl()
	}
}
