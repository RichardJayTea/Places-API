package handlers

import (
	"github.com/richardjaytea/teago/internal/app/teago/user"
	"net/http"
)

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var u user.User

	if err := decode(r, &u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	if err := u.CreateUser(h.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, u)
}
