package main

import (
	"log"
	"net/http"

	"github.com/Moritisimor/shawty/internal/handlers"
	"github.com/Moritisimor/shawty/internal/repo"
)

func main() {
	dbPath := "shawty.db"
	repo, err := repo.New(dbPath)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	http.HandleFunc("GET /api/status", handlers.StatusHandler)
	http.HandleFunc("GET /api/link/{link}", handlers.RedirectHandler(repo))
	http.HandleFunc("POST /api/alias", handlers.PostAliasHandler(repo))

	http.ListenAndServe(":8080", nil)
}
