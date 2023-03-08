package page

import (
	"html/template"
	"net/http"
)

func SearchPage(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./source/templates/research.html")

	if r.FormValue("search") != "" {
		if r.FormValue("search") == "all" {
			bdd.Show = bdd.Show[:0]
			for i := 0; i < len(bdd.Data); i++ {
				bdd.Show = append(bdd.Show, bdd.Data[i])
			}
		} else {
			bdd.Show = bdd.Show[0:0]
			for i := 0; i < 52; i++ {
				// name
				if bdd.Data[i].Name == r.FormValue("search") {
					bdd.Show = append(bdd.Show, bdd.Data[i])
				}

				// members
				for j := 0; j < len(bdd.Data[i].Members); j++ {
					if bdd.Data[i].Members[j] == r.FormValue("search") {
						bdd.Show = append(bdd.Show, bdd.Data[i])
					}
				}
			}
		}
	}

	err := page.ExecuteTemplate(w, "research.html", bdd)
	if err != nil {
		return
	}
}
