package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("http://localhost" + port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
