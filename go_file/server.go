package main

import (
	"fmt"
	"groupie-t/go_file/page"
	"net/http"
)

const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("source/"))

	page.Variable()

	http.HandleFunc("/home", page.HomePage)
	http.HandleFunc("/artist", page.ArtistPage)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("http://localhost" + port + "/home")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
