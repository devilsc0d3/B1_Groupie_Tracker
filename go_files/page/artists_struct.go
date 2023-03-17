package page

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"groupie-t/go_files/struct"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var bdd structure.Base
var ArrayGender [][]string

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
		"Group", "Creation date", "Location", "Dates", "No results for your search...", "Group, member, location, concert date", "Number of members", "First album date",
		"Since", "First album on", "Members", "Concerts", "Map", "Suggestions",
		"The page you were looking for cannot be found.", "Please go to the MusicMinder+ home page by clicking the button below", "First Album", "Gender"}
	bdd.Language.Fr = []string{"Accueil", "Recherche", "Langue", "envoyer",
		"Groupe", "Date de création", "Lieu", "Dates", "Aucun résultat pour votre recherche...", "Groupe, membre, lieu, date de concert", "Nombre de membres", "Date du premier album",
		"Depuis", "Premier album le", "Membres", "Concerts", "Carte", "Suggestions",
		"La page que vous recherchez est introuvable.", "Veuillez vous rendre sur la page d'accueil de MusicMinder+ en cliquant sur le bouton ci-dessous", "Premier Album", "Genre"}
	bdd.Language.Es = []string{"Inicio", "Búsqueda", "Idioma", "enviar",
		"Grupo", "Fecha de creación", "Ubicación", "Fechas", "No hay resultados para su búsqueda...", "Grupo, miembro, lugar, fecha del concierto", "Número de miembros", "Fecha del primer álbum",
		"Desde", "Primer álbum el", "Miembros", "Conciertos", "Mapa", "Sugerencias",
		"La página que buscaba no se encuentra.", "Vaya a la página de inicio de MusicMinder+ haciendo clic en el botón siguiente", "Primer álbum", "Género"}
	bdd.Language.Ge = []string{"Heimat", "Forschung", "Sprache", "senden",
		"Gruppe", "Erstellungsdatum", "Ort", "Daten", "Keine Ergebnisse für Ihre Suche...", "Gruppe, Mitglied, Ort, Konzertdatum", "Anzahl der Mitglieder", "Datum des ersten Albums",
		"Seit", "Erstes Album auf", "Mitglieder", "Konzerte", "Karte", "Vorschläge",
		"Die von Ihnen gesuchte Seite kann nicht gefunden werden.", "Bitte gehen Sie zur MusicMinder+ Homepage, indem Sie auf die folgende Schaltfläche klicken", "Erstes Album", "Geschlecht"}

	bdd.Language.CurrentLang = bdd.Language.En
	ReadFile()
	ReadFileList()
	GetApi("https://groupietrackers.herokuapp.com/api/artists", &bdd.Data)
	GetApi("https://groupietrackers.herokuapp.com/api/artists", &bdd.Show)
	for i := 0; i < len(bdd.Data); i++ {
		GetApi("https://groupietrackers.herokuapp.com/api/locations/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].Locations)
		GetApi("https://groupietrackers.herokuapp.com/api/dates/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].ConcertDates)
		GetApi("https://groupietrackers.herokuapp.com/api/relation/"+strconv.FormatInt(int64(i+1), 10), &bdd.Data[i].Relations)
	}

	for i := 0; i < 52; i++ {
		bdd.Data[i].Language = bdd.Language
		bdd.Data[i].Gender = ArrayGender[i]

	}
	TriGenre(bdd.DataGenre)
	//-----------------------------------//
	//for i := 0; i < 52; i++ {
	//	Spotify(bdd.Data[i].Name, i)
	//}
	//write()
	//-----------------------------------//
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 52; i++ {
		listShuffle := generateRandomList()
		for o := 0; o < len(listShuffle); o++ {
			bdd.Data[i].Suggestion = append(bdd.Data[i].Suggestion, bdd.Data[listShuffle[o]])
		}
	}

	for i := 0; i < len(bdd.Data); i++ {
		artist := &bdd.Data[i]
		http.HandleFunc("/"+strconv.FormatInt(artist.ID, 10), createHandlerFunction(artist))
	}
	http.HandleFunc("/research", SearchPage)
}

func TriGenre(array map[string][]structure.Artists) {
	temp := map[string][]structure.Artists{}
	for key, value := range array {
		if len(value) > 4 {
			temp[key] = value
		}
	}
	bdd.ShowGenre = temp
}

func write(file string) {
	b, _ := json.Marshal(gender)
	save, _ := os.Create(file)
	_, err := save.Write(b)
	if err != nil {
		return
	}
}

func ReadFile() {
	content, _ := os.ReadFile("../save.txt")
	err1 := json.Unmarshal(content, &bdd.DataGenre)
	if err1 != nil {
		fmt.Print(err1, "error")
		return
	}
}

func ReadFileList() {
	content, _ := os.ReadFile("SaveGender.txt")
	err1 := json.Unmarshal(content, &ArrayGender)
	if err1 != nil {
		fmt.Print(err1, "error")
		return
	}
}

func generateRandomList() []int {
	nums := make([]int, 30)
	for i := 0; i < 30; i++ {
		randomNum := rand.Intn(52)
		for IsPresent(nums, randomNum) {
			randomNum = rand.Intn(52)
		}
		nums[i] = randomNum
	}

	return nums
}

func IsPresent(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

var groups = map[string][]structure.Artists{}
var gender [][]string

func Spotify(group string, i int) {
	// create a config client for authenticator
	config := &clientcredentials.Config{
		ClientID:     "317c8b50b7974d9ea372963d712069fd",
		ClientSecret: "fbbb467fdfb6474f92719ab754b55fa4",
		TokenURL:     spotify.TokenURL,
	}
	// get a client authenticator
	token, err := config.Token(context.Background())
	if err != nil {
	}
	client := spotify.Authenticator{}.NewClient(token)

	// search artist on Spotify
	results, err := client.Search(group, spotify.SearchTypeArtist)
	if err != nil {
		fmt.Print(err)
	}

	artist := results.Artists.Artists[0]

	// get CATEGORIES
	fullArtist, err := client.GetArtist(artist.ID)
	if err != nil {
		fmt.Print(err)
	}

	//gender = append(gender, fullArtist.Genres)

	for _, category := range fullArtist.Genres {
		groups[category] = append(groups[category], bdd.Data[i])
	}

}

func createHandlerFunction(artist *structure.Artists) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ArtistPage(w, r, *artist)
	}
}
