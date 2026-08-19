package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strconv"

	"github.com/Moritisimor/shawty/internal/handlers"
	"github.com/Moritisimor/shawty/internal/helpers"
	"github.com/Moritisimor/shawty/internal/middleware"
	"github.com/Moritisimor/shawty/internal/repo"
)

//go:embed site
var staticFiles embed.FS

func main() {
	dbPath := helpers.GetEnvOr("DB_PATH", "shawty.db")
	address := helpers.GetEnvOr("ADDRESS", "0.0.0.0")
	port := helpers.GetEnvOr("PORT", "8080")
	sleepTimeMins := 1

	deleteAfterHours, err := strconv.ParseInt(helpers.GetEnvOr("DELETE_AFTER_HOURS", "24"), 10, 64)
	if err != nil {
		log.Fatalf("Error while parsing DELETE_AFTER_HOURS envvar: %s\n", err.Error())
	}

	log.Printf("Using database '%s'\n", dbPath)
	repo, err := repo.New(dbPath, deleteAfterHours)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	log.Printf("Connecting to database and migration successful\n")
	defer repo.Close()

	http.Handle("GET /api/status", middleware.Logging(http.HandlerFunc(handlers.StatusHandler)))
	http.Handle("GET /link/{link}", middleware.Logging(handlers.RedirectHandler(repo)))
	http.Handle("POST /api/alias", middleware.Logging(handlers.PostAliasHandler(repo)))

	siteFS, err := fs.Sub(staticFiles, "site")
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	http.Handle("GET /", middleware.Logging(http.FileServer(http.FS(siteFS))))

	helpers.StartCleanerRoutine(repo, sleepTimeMins)

	log.Printf("Listening on http://%s:%s\n", address, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil); err != nil {
		log.Fatalf("Error while listening: %s\n", err.Error())
	}
}
