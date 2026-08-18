package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		marshaledJSON, _ := json.Marshal(map[string]string{
			"status": "UP",
		})
		
		w.Write(marshaledJSON)
	})

	http.ListenAndServe(":8080", nil)
}
