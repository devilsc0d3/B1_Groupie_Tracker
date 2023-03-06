package page

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./source/templates/home.html")
	err := page.ExecuteTemplate(w, "home.html", bdd)
	if err != nil {
		return
	}

}
