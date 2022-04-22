package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	server     *http.Server
	ctx        context.Context
	signalChan chan os.Signal
	cancel     context.CancelFunc
}

// Create a new Server struct
func New(addr string, handler http.Handler) Server {
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for SIGHUP, SIGINT, SIGTERM and pass them in the channel
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	return Server{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		ctx:        ctx,
		signalChan: sig,
		cancel:     cancel,
	}
}

func (self *Server) ListenAndServe() {
	// Catch SIGINT and SIGTERM then attempt graceful server shutdown
	go func() {
		<-self.signalChan

		if err := self.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	// Run the server
	err := self.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for it to be done before exiting
	<-self.ctx.Done()
}

func (self *Server) Shutdown() error {
	shutdownCtx, cancel := context.WithTimeout(self.ctx, 30*time.Second)
	defer cancel()

	// Tell ListenAndServe that it can shutdown once this function exits
	defer self.cancel()

	// Crucify the server if it doesn't shutdown in 30 seconds
	go func() {
		<-shutdownCtx.Done()
		if shutdownCtx.Err() == context.DeadlineExceeded {
			log.Fatal("graceful shutdown timed out... forcing exit")
		}
	}()

	// Attempt the actual shutdown; returns error
	return self.server.Shutdown(shutdownCtx)
}
