package main

import (
	"context"
	"go-service/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "users-api ", log.LstdFlags)

	// create the handlers
	uh := handlers.NewUsers(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", uh)

	// create a new server
	s := http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  140 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	s.Shutdown(ctx)
}
