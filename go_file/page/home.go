package page

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Templates/" + "home" + ".html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}
