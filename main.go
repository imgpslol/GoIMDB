package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Tempoarily building movie list manually will come back to this later and build a database and connect to IMDB API to get movie list and details, tags and reviews.
type Movie struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World")

	// Create a handler function to serve the HTML template to index.html
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Remove later temporarily building movie list manually will come back to this later and build a database and connect to IMDB API to get movie list and details, tags and reviews.
		movies := map[string][]Movie{
			"Movies": {
				{Title: "Inception", Director: "Christopher Nolan"},
				{Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski"},
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
			},
		}
		tmpl.Execute(w, movies)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
