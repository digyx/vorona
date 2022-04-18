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

func New(addr string, handler http.Handler) Server {
	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	server := Server{
		server: &http.Server{
			Addr:    "0.0.0.0:8080",
			Handler: handler,
		},
		ctx:        ctx,
		signalChan: sig,
		cancel:     cancel,
	}

	return server
}

func (self *Server) ListenAndServe() {
	// Catch SIGINT and SIGTERM then attempt graceful server shutdown
	go func() {
		<-self.signalChan
		err := self.Shutdown()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Run the server
	err := self.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-self.ctx.Done()
}

func (self *Server) Shutdown() error {
	shutdownCtx, cancel := context.WithTimeout(self.ctx, 30*time.Second)
	defer cancel()

	go func() {
		<-shutdownCtx.Done()
		if shutdownCtx.Err() == context.DeadlineExceeded {
			log.Fatal("graceful shutdown timed out... forcing exit")
		}
	}()

	err := self.server.Shutdown(shutdownCtx)
	if err != nil {
		return err
	}

	self.cancel()

	return nil
}
