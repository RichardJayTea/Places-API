package handlers

import (
	"github.com/gorilla/mux"
	"github.com/richardjaytea/teago/cmd/teago/app/requests"
	"github.com/richardjaytea/teago/internal/app/teago/user"
	"net/http"
	"strconv"
)

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var d requests.CreateUser

	if err := decode(r, &d); err != nil {
		respondWithBadRequest(w)
		return
	}
	defer r.Body.Close()

	u, err := user.CreateUser(h.DB, d)
	if err != nil {
		respondWithInternalServerError(w, err)
	}

	respondWithJSON(w, http.StatusOK, u)
}

func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	// Put in middleware
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var d map[string]interface{}
	if err := decode(r, &d); err != nil {
		respondWithBadRequest(w)
	}

	u, err := user.UpdateUserByID(h.DB, id, d)
	if err != nil {
		respondWithInternalServerError(w, err)
	}

	respondWithJSON(w, http.StatusOK, u)
}
