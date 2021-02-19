package handlers

import (
	"github.com/richardjaytea/teago/internal/app/teago/place"
	"net/http"
)

func (h *Handler) HandleCreatePlace(w http.ResponseWriter, r *http.Request) {
	var d map[string]interface{}

	if err := decode(r, &d); err != nil {
		respondWithBadRequest(w)
		return
	}
	defer r.Body.Close()

	u, err := place.CreatePlace(h.DB, d)
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}
