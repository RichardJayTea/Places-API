package handlers

import (
	"github.com/richardjaytea/teago/internal/app/teago/checkin"
	"net/http"
)

func (h *Handler) HandleCreateCheckin(w http.ResponseWriter, r *http.Request) {
	var d map[string]interface{}

	if err := decode(r, &d); err != nil {
		respondWithBadRequest(w)
	}

	if err := checkin.CreateCheckin(h.DB, d); err != nil {
		respondWithInternalServerError(w, err)
	}
}
