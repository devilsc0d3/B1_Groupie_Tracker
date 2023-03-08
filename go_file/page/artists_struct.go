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
	Locations    Locations
	ConcertDates Dates
	Relations    string `json:"relations"`
}

type Base struct {
	Nbr    int
	Data   []Artists
	Show   []Artists
	SetImg []string
}

type Locations struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}

var bdd Base
var dates Dates

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
	for i := 0; i < len(bdd.Data); i++ {
		GetApi("https://groupietrackers.herokuapp.com/api/locations/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].Locations)
		GetApi("https://groupietrackers.herokuapp.com/api/dates/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].ConcertDates)
	}
	bdd.SetImg = []string{"ursula.png", "bob.webp", "err_404_8.webp"}
	//bdd.SetImg = append(bdd.SetImg, "ursula.png")
	//bdd.SetImg = append(bdd.SetImg, "bob.webp")
	for i := 0; i < len(bdd.Data); i++ {
		artist := bdd.Data[i]
		http.HandleFunc("/"+strconv.FormatInt(bdd.Data[i].ID, 10), func(w http.ResponseWriter, r *http.Request) {
			ArtistPage(w, r, artist)
		})
	}
	http.HandleFunc("/research", SearchPage)
}
