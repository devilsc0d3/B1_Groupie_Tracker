package page

import (
	"html/template"
	"math/rand"
	"net/http"
)

func Error(w http.ResponseWriter, _ *http.Request) {
	//TODO : ester eggs
	page, _ := template.ParseFiles("./source/templates/404.html")
	bdd.Nbr = RandomInt()
	err := page.ExecuteTemplate(w, "404.html", bdd)
	if err != nil {
		return
	}
}

func RandomInt() int {
	return rand.Intn(11)
}
