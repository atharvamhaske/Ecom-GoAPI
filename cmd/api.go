package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type application struct {
	config config
}

//run
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	
	//middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	
	// http.ListenAndServe(":8080", r)
	return r
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string //domain state name
}
