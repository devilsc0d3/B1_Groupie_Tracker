package page

import (
	"html/template"
	"math/rand"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./source/templates/404.html")
	bdd.Nbr = RandomInt()
	err := page.ExecuteTemplate(w, "404.html", bdd)
	if err != nil {
		return
	}
	//TODO : 404 different
}

func RandomInt() int {
	return rand.Intn(3)
}
