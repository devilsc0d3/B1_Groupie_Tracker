package page

import (
	"html/template"
	"math/rand"
	"net/http"
)

func Error(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("../source/templates/404.html")
	bdd.Nbr = RandomInt()
	err := page.ExecuteTemplate(w, "404.html", bdd)
	if err != nil {
		return
	}
}

func RandomInt() int {
	return rand.Intn(11)
}

func Test(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("../source/templates/index.html")
	err := page.ExecuteTemplate(w, "index.html", bdd)
	if err != nil {
		return
	}
}
