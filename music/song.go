package music

type Song struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Preview  string `json:"preview"`
	Mp3     string `json:"file"`
	Duration int    `json:"duration"`
	Rank     int    `json:"rank"`
	Artist   Artist
}
