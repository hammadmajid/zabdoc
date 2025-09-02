package main

import (
	"fmt"
	"net/http"
	"time"

	http2 "github.com/hammadmajid/zabcover/internal/api/http"
	"github.com/hammadmajid/zabcover/internal/app"
)

func main() {
	port := 8080
	zabcover, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	r := http2.SetupRoutes(zabcover)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Minute,
	}

	zabcover.Logger.Println(fmt.Sprintf("Listening on http://localhost:%d", port))

	err = server.ListenAndServe()
	if err != nil {
		zabcover.Logger.Fatalln(err)
	}
}
