package app

import "github.com/richardjaytea/teago/cmd/teago/app/handlers"

func (s *Server) routes(h handlers.Handler) {
	s.Router.HandleFunc("/user", h.HandleCreateUser).Methods("POST")
}
