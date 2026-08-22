package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

// Temporarily building movie list manually.
// Later this will be replaced with a database and TMDB API integration.
type Movie struct {
	Title    string
	Director string
}

func main() {

	// =========================
	// Load environment variables
	// =========================

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Loading GoIMDB...")

	// =========================
	// Homepage handler
	// =========================

	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(
			template.ParseFiles("index.html"),
		)

		// Temporarily building movie list manually.
		// Later this will be replaced with a database and TMDB API integration.
		movies := map[string][]Movie{
			"Movies": {
				{Title: "Inception", Director: "Christopher Nolan"},
				{Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski"},
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
			},
		}

		tmpl.Execute(w, movies)
	}

	// =========================
	// Add movie handler
	// =========================

	h2 := func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(1 * time.Second)

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		// Parses the index.html file and executes the
		// "movie-list-element" template block.
		tmpl := template.Must(
			template.ParseFiles("index.html"),
		)

		tmpl.ExecuteTemplate(
			w,
			"movie-list-element",
			Movie{
				Title:    title,
				Director: director,
			},
		)
	}

	// =========================
	// Routes
	// =========================

	http.HandleFunc("/", h1)

	http.HandleFunc("/add-movie/", h2)

	// Temporary TMDB API test
	http.HandleFunc("/tmdb-test", testTMDB)

	// =========================
	// Start server
	// =========================

	log.Println("GoIMDB running on http://localhost:8080")

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
