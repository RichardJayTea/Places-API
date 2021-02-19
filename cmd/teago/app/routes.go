package app

import "github.com/richardjaytea/teago/cmd/teago/app/handlers"

func (s *Server) routes(h handlers.Handler) {
	s.Router.HandleFunc("/user", h.HandleCreateUser).Methods("POST")
	s.Router.HandleFunc("/user/{id:[0-9]+}", h.HandleUpdateUser).Methods("PUT")

	s.Router.HandleFunc("/place", h.HandleCreatePlace).Methods("POST")
}
