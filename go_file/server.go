package main

import (
	"fmt"
	"groupie-t/go_file/page"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/home", page.HomePage)
	fmt.Println("http://localhost" + port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
