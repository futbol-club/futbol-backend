package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: @InProgress make Router struct.

func GetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// TODO: @InProgress create a method to handle this routes.
	router.HandleFunc("/", IndexRoute)

	return router
}

func IndexRoute(writer http.ResponseWriter, request *http.Request) {
	logger(request)
	fmt.Fprintf(writer, "Hello, Sailor: %q\n", html.EscapeString(request.URL.Path))
}
