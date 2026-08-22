package main

import (
	"fmt"
	"net/http"
	"os"
)

func testTMDB(w http.ResponseWriter, r *http.Request) {

	// Create a request to TMDB.
	// Movie ID 550 is Fight Club.
	req, err := http.NewRequest(
		"GET",
		"https://api.themoviedb.org/3/movie/550",
		nil,
	)

	if err != nil {
		http.Error(
			w,
			"Failed to create request",
			http.StatusInternalServerError,
		)

		return
	}

	// Get the TMDB token from the environment.
	tmdbToken := os.Getenv("TMDB_TOKEN")

	// Add the token to the request.
	req.Header.Set(
		"Authorization",
		"Bearer "+tmdbToken,
	)

	// Send the request to TMDB.
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		http.Error(
			w,
			"Failed to contact TMDB",
			http.StatusInternalServerError,
		)

		return
	}

	defer response.Body.Close()

	// Display the response status.
	fmt.Fprintf(
		w,
		"TMDB returned status: %s",
		response.Status,
	)
}
