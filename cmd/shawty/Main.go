package main

import (
	"net/http"

	"github.com/Moritisimor/shawty/internal/handlers"
)

func main() {
	http.HandleFunc("/api/status", handlers.StatusHandler)

	http.ListenAndServe(":8080", nil)
}
