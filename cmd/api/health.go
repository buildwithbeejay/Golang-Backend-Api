package main

import (
	"log"
	"net/http"
)


func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte("ok")); err != nil {
	// Log the error; response already started so can't change status
		log.Printf("health check write failed: %v", err)
	}
}	