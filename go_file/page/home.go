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
	for i := 0; i < 52; i++ {
		Spotify(bdd.Data[i].Name, i)
	}

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", 303)
	}

	page, _ := template.ParseFiles("./source/templates/home.html")
	err := page.ExecuteTemplate(w, "home.html", bdd)
	if err != nil {
		return
	}

}

func Spotify(group string, i int) {
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

	// Sélectionner le premier artiste trouvé
	artist := results.Artists.Artists[0]

	// Récupérer les catégories de musique de l'artiste
	fullartist, err := client.GetArtist(artist.ID)
	if err != nil {
		fmt.Print(err)
	}

	groups := map[string][]Artists{}

	// Ajouter l'artiste au dico
	for _, category := range fullartist.Genres {
		groups[category] = append(groups[category], bdd.Data[i])
	}

	// test affiche categories
	for key := range groups {
		fmt.Println(key)
	}
}
