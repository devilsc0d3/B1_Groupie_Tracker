package page

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./source/templates/home.html")

	if r.FormValue("search") != "" {
		if r.FormValue("search") == "all" {
			bdd.Show = bdd.Show[:0]
			for i := 0; i < len(bdd.Data); i++ {
				bdd.Show = append(bdd.Show, bdd.Data[i])
			}
		} else {
			for i := 0; i < 52; i++ {
				if bdd.Data[i].Name == r.FormValue("search") {
					bdd.Show = bdd.Show[0:0]
					bdd.Show = append(bdd.Show, bdd.Data[i])
				}
			}
		}
	}
	err := page.ExecuteTemplate(w, "home.html", bdd)
	if err != nil {
		return
	}

}
