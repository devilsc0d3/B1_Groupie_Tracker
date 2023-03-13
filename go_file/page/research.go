package page

import (
	"html/template"
	"net/http"
	"strconv"
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
				if ToLower(bdd.Data[i].Name) == ToLower(r.FormValue("search")) {
					bdd.Show = append(bdd.Show, bdd.Data[i])
				}

				// members
				for j := 0; j < len(bdd.Data[i].Members); j++ {
					if ToLower(bdd.Data[i].Members[j]) == ToLower(r.FormValue("search")) {
						bdd.Show = append(bdd.Show, bdd.Data[i])
					}
				}

				//locations
				for j := 0; j < len(bdd.Data[i].Locations.Locations); j++ {
					if ToLower(bdd.Data[i].Locations.Locations[j]) == ToLower(r.FormValue("search")) {
						bdd.Show = append(bdd.Show, bdd.Data[i])
					}
				}

				//dates
				for j := 0; j < len(bdd.Data[i].ConcertDates.Dates); j++ {
					if bdd.Data[i].ConcertDates.Dates[j] == r.FormValue("search") {
						bdd.Show = append(bdd.Show, bdd.Data[i])
					}
				}

				//creation dates
				if strconv.FormatInt(bdd.Data[i].CreationDate, 10) == r.FormValue("search") {
					bdd.Show = append(bdd.Show, bdd.Data[i])
				}
			}
		}
	}

	if r.FormValue("language") == "en" {
		bdd.Language.CurrentLang = bdd.Language.En
	} else if r.FormValue("language") == "fr" {
		bdd.Language.CurrentLang = bdd.Language.Fr
	} else if r.FormValue("language") == "es" {
		bdd.Language.CurrentLang = bdd.Language.Es
	} else if r.FormValue("language") == "ge" {
		bdd.Language.CurrentLang = bdd.Language.Ge
	}

	err := page.ExecuteTemplate(w, "research.html", bdd)
	if err != nil {
		return
	}
}

func ToLower(s string) string {
	tr := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			tr += string(rune(s[i] + 32))
		} else {
			tr += string(rune(s[i]))
		}
	}
	return tr
}
