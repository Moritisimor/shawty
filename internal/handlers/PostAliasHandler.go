package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Moritisimor/shawty/internal/helpers"
	"github.com/Moritisimor/shawty/internal/models"
	"github.com/Moritisimor/shawty/internal/repo"
	"github.com/ncruces/go-sqlite3"
)

func PostAliasHandler(repo *repo.URLAliasRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var dto models.URLAliasDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&dto); err != nil {
			helpers.SendJSON(w, http.StatusUnprocessableEntity, helpers.J{
				"error": "Bad JSON",
			})

			return
		}

		if !dto.IsValid() {
			helpers.SendJSON(w, http.StatusBadRequest, helpers.J{
				"error": "The JSON you sent has one or more missing/empty fields",
			})

			return
		}

		id, err := repo.PostURLAlias(dto, r.Context()) 
		if err != nil {
			if errors.Is(err, sqlite3.CONSTRAINT) {
				helpers.SendJSON(w, http.StatusConflict, helpers.J{
					"error": "Alias already taken",
				})
			} else {
				helpers.SendJSON(w, http.StatusInternalServerError, helpers.J{
					"error": "Error while writing to database",
				})
			}

			return
		}

		helpers.SendJSON(w, http.StatusOK, helpers.J{
			"id": id,
		})
	}
}
