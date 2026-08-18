package handlers

import (
	"net/http"

	"github.com/Moritisimor/shawty/internal/helpers"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendJSON(w, http.StatusOK, helpers.J{
		"status": "UP",
	})
}
