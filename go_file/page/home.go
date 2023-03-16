package page

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("language") != "" {
		if r.FormValue("language") == "en" {
			bdd.Language.CurrentLang = bdd.Language.En
		} else if r.FormValue("language") == "fr" {
			bdd.Language.CurrentLang = bdd.Language.Fr
		} else if r.FormValue("language") == "es" {
			bdd.Language.CurrentLang = bdd.Language.Es
		} else if r.FormValue("language") == "ge" {
			bdd.Language.CurrentLang = bdd.Language.Ge
		}
		for i := 0; i < 52; i++ {
			bdd.Data[i].Language.CurrentLang = bdd.Language.CurrentLang
		}
	}


	//// test affiche categories
	//for key, valeur := range groups {
	//	fmt.Println(key, ":")
	//	for i := 0; i < len(valeur); i++ {
	//		fmt.Println(valeur[i].Name)
	//	}
	//}

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", 303)
	}

	page, _ := template.ParseFiles("../source/templates/home.html")
	err := page.ExecuteTemplate(w, "home.html", bdd)
	if err != nil {
		return
	}

}

var groups = map[string][]Artists{}

func Spotify(group string, i int) {
	// Créer une configuration client credentials pour l'authentification OAuth2
	config := &clientcredentials.Config{
		ClientID:     "317c8b50b7974d9ea372963d712069fd",
		ClientSecret: "fbbb467fdfb6474f92719ab754b55fa4",
		TokenURL:     spotify.TokenURL,
	}
	// Obtenir un client authentifié avec la configuration client credentials
	token, err := config.Token(context.Background())
	if err != nil {
	}
	client := spotify.Authenticator{}.NewClient(token)

	// Rechercher l'artiste sur Spotify
	results, err := client.Search(group, spotify.SearchTypeArtist)
	if err != nil {
		fmt.Print(err)
	}

	// Sélectionner le premier artiste trouvé
	artist := results.Artists.Artists[0]

	// Récupérer les catégories de musique de l'artiste
	fullartist, err := client.GetArtist(artist.ID)
	if err != nil {
		fmt.Print(err)
	}

	// Ajouter l'artiste au dico
	for _, category := range fullartist.Genres {
		groups[category] = append(groups[category], bdd.Data[i])
	}

}
