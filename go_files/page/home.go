package page

import (
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

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", 303)
	}

	page, _ := template.ParseFiles("../source/templates/home.html")
	err := page.ExecuteTemplate(w, "home.html", bdd)
	if err != nil {
		return
	}

}
