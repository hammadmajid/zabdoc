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

	http2 "zabdoc/internal/api/http"
	"zabdoc/internal/app"
	"zabdoc/internal/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("env PORT is not defined")
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	handler := http2.NewHandler(logger)

	app := app.NewApp(logger, handler)

	r := router.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Minute,
	}

	app.Logger.Printf("Listening on http://localhost:%s", port)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.Logger.Fatalf("Server forced to shutdown: %v", err)
	}

	app.Logger.Println("Server exited")
}
