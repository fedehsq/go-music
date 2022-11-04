package music

type Song struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Preview  string `json:"preview"`
	Mp3      string `json:"file"`
	Duration int    `json:"duration"`
	Rank     int    `json:"rank"`
	Url      string  
	Artist   Artist
}
