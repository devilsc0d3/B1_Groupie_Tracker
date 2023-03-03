package page

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./source/templates/home.html")
	fmt.Print(bdd)
	if r.FormValue("search") != "" {
		if r.FormValue("search") == "all" {
			bdd.Show = bdd.Data
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
