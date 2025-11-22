package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	
	//good base middlewares stack
	r.Use(middleware.RequestID) //important for rate limiting
	r.Use(middleware.RealIP) //important for tracing and rate limiting and analytics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//set a timeout value on the request context, that will signal through
	//ctx.Done() that the request has time out and further processing
	//should stop
	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	
	// http.ListenAndServe(":8080", r)
	return r
}

//run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	
	//channel to catch OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	log.Printf("server started at %s", app.config.addr)
	
	//run the server in a goroutine (so it doesn't block)
	go func() {
		if err := srv.ListenAndServe();
		err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server could not listen: %v\n", err)
		}
	}()
	
	//block until a signal is recieved
	<- quit //pause the porgram until something is received on quit shorthand for _ = <- quit
	log.Println("shutting down server")
	
	//give server up to 5 secs to finish requests
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx);
	err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}
	
	log.Println("server stopped gracefully")
	return nil
}


type application struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string //domain state name
}
