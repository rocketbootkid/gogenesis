package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func generate(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/generate", generate)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
