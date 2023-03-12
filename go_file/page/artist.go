package page

import (
	"html/template"
	"net/http"
)

func ArtistPage(w http.ResponseWriter, _ *http.Request, data interface{}) {
	page, _ := template.ParseFiles("./source/templates/artist.html")
	err := page.ExecuteTemplate(w, "artist.html", data)
	if err != nil {
		return
	}
}
