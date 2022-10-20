package artist

import "encoding/json"

type Artist struct {
	Id int
	Name string
	Picture string
}

func FromJson(jsong []byte) Artist {
	var artist Artist
	json.Unmarshal(jsong, &artist)
	return artist
}