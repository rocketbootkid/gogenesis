package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func generate(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		rows, ok := r.URL.Query()["rows"]
		if !ok || len(rows[0]) < 1 {
			http.Error(w, "rows parameter missing", http.StatusBadRequest)
			return
		}
		numRows, strerr := strconv.Atoi(rows[0])
		if strerr != nil {
			http.Error(w, strerr.Error(), http.StatusInternalServerError)
			return
		}

		var cols []Column
		err := json.NewDecoder(r.Body).Decode(&cols)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Column: %+v", cols)

		m := make(map[string]string)

		m["Forename"] = "N:f"
		m["Surname"] = "N:s"

		data := DataSchema{numRows, m}
		json, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

	default:
		http.Error(w, "Only POST supported", http.StatusBadRequest)
		return
	}

}

type DataSchema struct {
	Rows    int
	Columns map[string]string
}

type Column struct {
	definition string
}

func handleRequests() {
	http.HandleFunc("/generate", generate)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
