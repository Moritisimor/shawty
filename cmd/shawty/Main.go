package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Moritisimor/shawty/internal/handlers"
	"github.com/Moritisimor/shawty/internal/helpers"
	"github.com/Moritisimor/shawty/internal/middleware"
	"github.com/Moritisimor/shawty/internal/repo"
)

func main() {
	dbPath := helpers.GetEnvOr("DB_PATH", "shawty.db")
	address := helpers.GetEnvOr("ADDRESS", "0.0.0.0")
	port := helpers.GetEnvOr("PORT", "8080")

	log.Printf("Using database '%s'\n", dbPath)
	repo, err := repo.New(dbPath)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	log.Printf("Connecting to database and migration successful\n")
	defer repo.Close()

	http.Handle("GET /api/status", middleware.Logging(handlers.StatusHandler))
	http.Handle("GET /link/{link}", middleware.Logging(handlers.RedirectHandler(repo)))
	http.Handle("POST /api/alias", middleware.Logging(handlers.PostAliasHandler(repo)))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Printf("Listening on http://%s:%s\n", address, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil); err != nil {
		log.Fatalf("Error while listening: %s\n", err.Error())
	}
}
