package page

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

type Base struct {
	Data []Artists
	Show []Artists
}

type Locations struct {
	ID        int64  `json:"id"`
	Locations string `json:"locations"`
}

type Welcome struct {
	Artists   []Base
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

var bdd Base
var welcome Welcome

func GetApi(url string, target interface{}) {
	response, err := http.Get(url)
	if err != nil {
		print("err")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(data, &target)
		if err != nil {
			return
		}
	}
}

func Variable() {
	GetApi("https://groupietrackers.herokuapp.com/api/artists", &bdd.Data)
	GetApi("https://groupietrackers.herokuapp.com/api/artists", &bdd.Show)
	GetApi("https://groupietrackers.herokuapp.com/api", &welcome)

	for i := 0; i < len(bdd.Data); i++ {
		artist := bdd.Data[i]
		http.HandleFunc("/"+strconv.FormatInt(bdd.Data[i].ID, 10), func(w http.ResponseWriter, r *http.Request) {
			ArtistPage(w, r, artist)
		})
	}
}
