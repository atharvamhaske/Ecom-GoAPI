package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

//run
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	
	//good base middlewares stack
	r.Use(middleware.RequestID) //important for rate limiting
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	
	//set a timeout value on the request context, that will signal through
	//ctx.Done() that the request has time out and further processing
	//should stop
	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
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
