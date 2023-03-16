package page

import (
	"html/template"
	"net/http"
)

func ArtistPage(w http.ResponseWriter, r *http.Request, data interface{}) {
	page, _ := template.ParseFiles("../source/templates/artist.html")
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
			bdd.Data[i].Language = bdd.Language
		}
	}
	err := page.ExecuteTemplate(w, "artist.html", data)
	if err != nil {
		return
	}
}
