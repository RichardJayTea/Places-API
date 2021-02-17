package app

import "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/cmd/teago/app/handlers"

func (s *Server) routes(h handlers.Handler) {
	s.Router.HandleFunc("/user", h.HandleCreateUser).Methods("POST")
}
