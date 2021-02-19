package handlers

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func respondWithBadRequest(w http.ResponseWriter) {
	respondWithError(w, http.StatusBadRequest, "Invalid request body")
}

func respondWithInternalServerError(w http.ResponseWriter, err error) {
	respondWithError(w, http.StatusInternalServerError, err.Error())
}
