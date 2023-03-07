package page

import (
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./source/templates/404.html")
	err := page.ExecuteTemplate(w, "404.html", bdd)
	if err != nil {
		return
	}
	//TODO : 404 different
}

func RandomInt() int {
	return 5
}
