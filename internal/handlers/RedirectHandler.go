package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/Moritisimor/shawty/internal/helpers"
	"github.com/Moritisimor/shawty/internal/repo"
)

// /api/link/{link}
func RedirectHandler(repo *repo.URLAliasRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		short := r.PathValue("link")
		url, err := repo.GetURLByAlias(short, r.Context())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				helpers.SendJSON(w, http.StatusNotFound, helpers.J{
					"error": "No such alias",
				})
			} else {
				log.Printf("Error while querying: %s\n", err.Error())
				helpers.SendJSON(w, http.StatusInternalServerError, helpers.J{
					"error": "Error while querying database",
				})
			}

			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	}
}
