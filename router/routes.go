package router

import (
	"net/http"

	"github.com/fkocharli/savejoe/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"StreamSeq",
		"GET",
		"/stream/{filename}/{sequence}",
		handlers.Play,
	},
	Route{
		"Stream",
		"GET",
		"/stream/{filename}/",
		handlers.Play,
	},
}
