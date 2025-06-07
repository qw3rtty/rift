package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// HTTPListener encapsulates the configuration and lifecycle of an HTTP server.
type HTTPListener struct {
	Addr              string        // Address to listen on, e.g., ":8080"
	Handler           http.Handler  // Custom handler for HTTP requests
	ReadTimeout       time.Duration // Timeout for reading the entire request
	WriteTimeout      time.Duration // Timeout for writing the response
	IdleTimeout       time.Duration // Timeout for idle connections
	TLSConfig         *tls.Config   // Optional TLS configuration
	GracefulShutdown  time.Duration // Graceful shutdown timeout
	ShutdownSignal    os.Signal     // Signal to initiate shutdown (default: os.Interrupt)
	Listener          net.Listener  // Underlying network listener
	Server            *http.Server  // HTTP server instance
}

// NewHTTPListener creates a new HTTPListener with the provided configuration.
func NewHTTPListener(addr string, handler http.Handler) *HTTPListener {
	return &HTTPListener{
		Addr:             addr,
		Handler:          handler,
		ReadTimeout:      10 * time.Second,
		WriteTimeout:     10 * time.Second,
		IdleTimeout:      60 * time.Second,
		GracefulShutdown: 5 * time.Second,
		ShutdownSignal:   os.Interrupt,
	}
}

// Start initializes and starts the HTTP listener.
func (hl *HTTPListener) Start() error {
	var err error
	hl.Listener, err = net.Listen("tcp", hl.Addr)
	if err != nil {
		return fmt.Errorf("[!] failed to start listener: %w", err)
	}

	hl.Server = &http.Server{
		Addr:              hl.Addr,
		Handler:           hl.Handler,
		ReadTimeout:       hl.ReadTimeout,
		WriteTimeout:      hl.WriteTimeout,
		IdleTimeout:       hl.IdleTimeout,
		TLSConfig:         hl.TLSConfig,
	}

	go func() {
		if hl.TLSConfig != nil {
			if err := hl.Server.ServeTLS(hl.Listener, "", ""); err != nil && err != http.ErrServerClosed {
				fmt.Printf("[!] TLS server failed: %v\n", err)
			}
		} else {
			if err := hl.Server.Serve(hl.Listener); err != nil && err != http.ErrServerClosed {
				fmt.Printf("[!] HTTP server failed: %v\n", err)
			}
		}
	}()

	return nil
}

// Shutdown gracefully shuts down the HTTP listener.
func (hl *HTTPListener) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), hl.GracefulShutdown)
	defer cancel()

	if err := hl.Server.Shutdown(ctx); err != nil {
		return fmt.Errorf("[!] server shutdown failed: %w", err)
	}

	if err := hl.Listener.Close(); err != nil {
		return fmt.Errorf("[!] listener close failed: %w", err)
	}

	return nil
}

// WaitForShutdown blocks until the specified shutdown signal is received.
func (hl *HTTPListener) WaitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, hl.ShutdownSignal)
	<-stop
	fmt.Println("Received shutdown signal")
}

