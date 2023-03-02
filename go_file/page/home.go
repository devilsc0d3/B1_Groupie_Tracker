package page

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./source/templates/home.html")

	if r.FormValue("search") != "" {
		if r.FormValue("search") == "all" {
			bbd.Show = bbd.Data
		} else {
			for i := 0; i < 52; i++ {
				if bbd.Data[i].Name == r.FormValue("search") {
					bbd.Show = bbd.Show[0:0]
					bbd.Show = append(bbd.Show, bbd.Data[i])
				}
			}
		}
	}
	err := page.ExecuteTemplate(w, "home.html", bbd)
	if err != nil {
		return
	}

}
