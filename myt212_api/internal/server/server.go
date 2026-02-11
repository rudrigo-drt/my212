package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string) *Server {
	srv := &http.Server{
		Addr:         addr,
		Handler:      NewRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &Server{httpServer: srv}
}

func (s *Server) Start() {
	//Capture OS signals (CTRL+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Starting server on " + s.httpServer.Addr)

		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Awaits interruption signal
	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server Shutdown failed: %v", err)
	}

	log.Println("Server shutdown successfully")
}
