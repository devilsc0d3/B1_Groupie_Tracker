package groupie_tracker

import "encoding/json"

type Artist []Artists

func UnmarshalArtist(data []byte) (Artist, error) {
	var r Artist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Artist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Artists struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
