package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"test-verify-app/internal/server"
)

func gracefulShutdown(apiServer *server.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	shutdownDuration := 5 * time.Second
	if t := os.Getenv("SHUTDOWN_TIMEOUT"); t != "" {
		if i, err := strconv.Atoi(t); err == nil {
			shutdownDuration = time.Duration(i) * time.Second
		}
	}

	// The context is used to inform the server it has 5 seconds (or configured value) to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), shutdownDuration)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	if err := apiServer.Close(); err != nil {
		log.Printf("Server failed to close resources: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {

	server := server.NewServer()

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
