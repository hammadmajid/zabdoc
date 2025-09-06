package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hammadmajid/zabcover/internal/app"
	httpapi "github.com/hammadmajid/zabcover/internal/router"
)

func main() {
	port := 8080
	zabcover, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	r := httpapi.SetupRoutes(zabcover)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Minute,
	}

	zabcover.Logger.Printf("Listening on http://localhost:%d", port)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zabcover.Logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zabcover.Logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zabcover.Logger.Fatalf("Server forced to shutdown: %v", err)
	}

	zabcover.Logger.Println("Server exited")
}
