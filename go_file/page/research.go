package page

import (
	"fmt"
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
	//lang
	if r.FormValue("language") == "en" {
		bdd.Language.CurrentLang = bdd.Language.En
	} else if r.FormValue("language") == "fr" {
		bdd.Language.CurrentLang = bdd.Language.Fr
	} else if r.FormValue("language") == "es" {
		bdd.Language.CurrentLang = bdd.Language.Es
	} else if r.FormValue("language") == "ge" {
		bdd.Language.CurrentLang = bdd.Language.Ge
	}

	//filters
	var listeu []int
	if r.FormValue("checkbox1") == "on" {
		listeu = append(listeu, 1)
	}
	if r.FormValue("checkbox2") == "on" {
		listeu = append(listeu, 2)
	}
	if r.FormValue("checkbox3") == "on" {
		fmt.Println("Checkbox 3 selected")
	}
	if r.FormValue("checkbox4") == "on" {
		fmt.Println("Checkbox 4 selected")
	}
	if r.FormValue("checkbox5") == "on" {
		fmt.Println("Checkbox 5 selected")
	}
	if r.FormValue("checkbox6") == "on" {
		fmt.Println("Checkbox 6 selected")
	}
	if r.FormValue("checkbox7") == "on" {
		listeu = append(listeu, 7)
	}
	if r.FormValue("checkbox8") == "on" {
		listeu = append(listeu, 8)
	}
	if len(listeu) == 0 {
		listeu = []int{1, 2, 3, 4, 5, 6, 7, 8}
	}
	filters(listeu)

	err := page.ExecuteTemplate(w, "research.html", bdd)
	if err != nil {
		return
	}
}

// TODO : same input & function refactor
func filters(nbr []int) {
	var temp []Artists
	for _, s := range bdd.Show {
		for _, n := range nbr {
			if len(s.Members) == n {
				temp = append(temp, s)
				break
			}
		}
	}
	bdd.Show = temp
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
