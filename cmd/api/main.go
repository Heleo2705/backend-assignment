package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/backend-assignment/handlers"
	middlewares "example.com/backend-assignment/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	err := handlers.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	err = handlers.InitMeili()
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	mux := chi.NewRouter()
	v1Router := chi.NewRouter()
	v1Router.Use(middleware.Logger)
	v1Router.Use(middleware.Recoverer)
	// user apis
	v1Router.Group(func(r chi.Router) {
		r.Post("/auth/signup", handlers.CreateUser)
		r.Post("/auth/login", handlers.Login)

	})
	// notes apis
	v1Router.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/notes", handlers.CreateNote)
		r.Get("/notes", handlers.GetNotesForUser)
		r.Put("/notes",handlers.UpdateNote)
		r.Delete("/notes",handlers.DeleteNotes)
		
	})
	//sharing apis
	v1Router.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/notes/{id}/share",handlers.ShareNote)
		r.Get("/notes/share",handlers.GetSharedNotes)
	})
	//searching apis
	v1Router.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Get("/search",handlers.SearchNotes)
	})
	mux.Mount("/api", v1Router)
	fmt.Printf("Server is starting..")
	server := http.Server{Addr: ":3000", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
