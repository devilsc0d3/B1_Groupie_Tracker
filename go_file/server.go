package main

import (
	"fmt"
	"groupie-t/go_file/page"
	"net/http"
)

const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("../source/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	page.Variable()

	http.HandleFunc("/test", page.Test)
	http.HandleFunc("/", page.HomePage)
	http.HandleFunc("/404", page.Error)

	fmt.Println("http://localhost" + port + "/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
