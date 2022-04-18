package server

import (
	"fmt"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"
)

// Ensure that the right signals trigger a graceful shutdown
func TestGracefulShutdown(t *testing.T) {
	testSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	}

	for index, sig := range testSignals {
		t.Run(sig.String(), func(t *testing.T) {
			// Make sure the servers don't listen to the same port
			addr := fmt.Sprintf("127.0.0.1:999%d", index)
			server := New(addr, http.DefaultServeMux)

			go func() {
				server.ListenAndServe()
			}()

			// Give it a secone to actually start
			time.Sleep(500 * time.Millisecond)

			// Send the signal to the webserver
			server.signalChan <- sig
		})
	}
}
