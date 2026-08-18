package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type J map[string]any

func SendJSON(w http.ResponseWriter, status int, content any) {
	w.Header().Add("Content-Type", "application/json")
	marshaledJSON, err := json.Marshal(content)
	if err != nil {
		log.Printf("Error while marshaling JSON: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\":\"Server Could not marshal JSON\"}"))

		return
	}

	w.WriteHeader(status)
	w.Write(marshaledJSON)
}
