package main

import (
	"log"
	"net/http"

	"github.com/fkocharli/savejoe/router"
	"github.com/rs/cors"
)

// setup CORS middleware
func runMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

func main() {
	// create router and listen 8080
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", runMiddleware(router)))
}
