package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"test-docs-project/internal/server"
)

func gracefulShutdown(apiServer *server.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	slog.Info("shutting down gracefully, press Ctrl+C again to force")
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
		slog.Error("Server forced to shutdown with error", "error", err)
	}

	if err := apiServer.Close(); err != nil {
		slog.Error("Server failed to close resources", "error", err)
	}

	slog.Info("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	var handler slog.Handler
	if os.Getenv("APP_ENV") == "local" {
		handler = slog.NewTextHandler(os.Stdout, nil)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)

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
	slog.Info("Graceful shutdown complete.")
}
