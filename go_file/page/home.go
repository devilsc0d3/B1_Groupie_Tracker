package page

import (
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
