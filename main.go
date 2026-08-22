package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Tempoarily building movie list manually will come back to this later and build a database and connect to IMDB API to get movie list and details, tags and reviews.
type Movie struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Loading GoIMDB...")

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

	// handler function #2 - returns the template block with the newly added film, as an HTMX response (WIP)
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		// parses the index.html file and executes the "movie-list-element" template block with the newly added movie data, which is then sent back as an HTMX response to be inserted into the existing movie list on the page.
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "movie-list-element", Movie{Title: title, Director: director})
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-movie/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
