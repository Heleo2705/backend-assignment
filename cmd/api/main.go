package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	v1Router:= chi.NewRouter()
	v1Router.Use(middleware.Logger)
	v1Router.Use(middleware.Recoverer)
	// user apis
	v1Router.Group(func(r chi.Router) {})
}