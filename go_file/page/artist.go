package page

import (
	"html/template"
	"net/http"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Templates/" + "artist" + ".html")
	err := page.ExecuteTemplate(w, "artist.html", nil)
	if err != nil {
		return
	}
}
