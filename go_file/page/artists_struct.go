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
	Relations    Relation
}

type Language struct {
	En          []string
	Fr          []string
	Es          []string
	Ge          []string
	CurrentLang []string
}

type Base struct {
	Nbr      int
	Data     []Artists
	Show     []Artists
	Language Language
}

type Relation struct {
	ID       int64               `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
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
var rel Relation

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

	bdd.Language.En = []string{"Home", "Research", "Language", "send",
		"No results for your search...", "Group, member, location, concert date",
		"Since", "First album on", "Members", "Concerts", "Map", "Suggestions"}
	bdd.Language.Fr = []string{"Accueil", "Recherche", "Langue", "envoyer",
		"Aucun résultat pour votre recherche...",
		"Depuis", "Premier album le", "Membres", "Concerts", "Carte", "Suggestions"}
	bdd.Language.Es = []string{"Inicio", "Búsqueda", "Idioma", "enviar",
		"No hay resultados para su búsqueda...", "Group, member, location, concert date",
		"Desde", "Primer álbum el", "Miembros", "Conciertos", "Mapa", "Sugerencias"}
	bdd.Language.Ge = []string{"Heimat", "Forschung", "Sprache", "senden",
		"Keine Ergebnisse für Ihre Suche...", "Group, member, location, concert date",
		"Seit", "Erstes Album auf", "Mitglieder", "Konzerte", "Karte", "Vorschläge"}

	bdd.Language.CurrentLang = bdd.Language.En

	GetApi("https://groupietrackers.herokuapp.com/api/artists", &bdd.Data)
	GetApi("https://groupietrackers.herokuapp.com/api/artists", &bdd.Show)
	for i := 0; i < len(bdd.Data); i++ {
		GetApi("https://groupietrackers.herokuapp.com/api/locations/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].Locations)
		GetApi("https://groupietrackers.herokuapp.com/api/dates/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].ConcertDates)
		GetApi("https://groupietrackers.herokuapp.com/api/relation/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].Relations)
	}

	for i := 0; i < len(bdd.Data); i++ {
		artist := bdd.Data[i]
		http.HandleFunc("/"+strconv.FormatInt(bdd.Data[i].ID, 10), func(w http.ResponseWriter, r *http.Request) {
			ArtistPage(w, r, artist)
		})
	}
	http.HandleFunc("/research", SearchPage)
}
