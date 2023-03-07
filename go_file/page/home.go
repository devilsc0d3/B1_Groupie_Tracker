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
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", 303)
	}

	page, _ := template.ParseFiles("./source/templates/home.html")
	err := page.ExecuteTemplate(w, "home.html", bdd)
	if err != nil {
		return
	}

}

func Spotify(group string) {
	// Créer une configuration client credentials pour l'authentification OAuth2
	config := &clientcredentials.Config{
		ClientID:     "73b0824483b349a7bccb211d72dff641",
		ClientSecret: "104f04fdfb984e54a3655f740dc64c7c",
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

	// Vérifier si des artistes ont été trouvés TEST
	if len(results.Artists.Artists) == 0 {
		fmt.Println("Aucun artiste trouvé pour", group)
	}

	// Sélectionner le premier artiste trouvé
	artist := results.Artists.Artists[0]

	// Récupérer les catégories de musique de l'artiste
	fullartist, err := client.GetArtist(artist.ID)
	if err != nil {
		fmt.Print(err)
	}

	// Ajouter l'artiste à la liste en fonction de ses catégories de musique
	for _, category := range fullartist.Genres {
		switch category {
		case "jazz":
			fmt.Println(group, "est un artiste de jazz")
			// Ajouter l'artiste à la liste de jazz
		case "blues":
			fmt.Println(group, "est un artiste de blues")
			// Ajouter l'artiste à la liste de blues
		default:
			fmt.Println(group, "est un artiste de", category)
			// Ajouter l'artiste à une liste générique pour cette catégorie de musique
		}
	}
}
