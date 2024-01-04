package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mux := chi.NewRouter()
	v1Router := chi.NewRouter()
	v1Router.Use(middleware.Logger)
	v1Router.Use(middleware.Recoverer)
	// user apis
	v1Router.Group(func(r chi.Router) {})
	// notes apis
	v1Router.Group(func(r chi.Router) {})
	mux.Mount("/api", v1Router)
	fmt.Printf("Server is starting..")
	server := http.Server{Addr: "8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
