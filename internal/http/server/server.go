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
}

func New(addr string, handler http.Handler) (Server, context.CancelFunc) {
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
	}

	return server, cancel
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

	err := self.server.Shutdown(shutdownCtx)
	if err != nil {
		return err
	}

	return nil
}
