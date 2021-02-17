package app

import (
	"database/sql"
	"fmt"
	"github.com/richardjaytea/teago/cmd/teago/app/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Server struct holds server config and dependencies
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize DB and Router for Server struct
func (s *Server) Initialize(user, password, dbname string) error {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, dbname)

	var err error
	s.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	h := handlers.Handler{DB: s.DB}

	s.Router = mux.NewRouter()
	s.routes(h)

	return nil
}

// Run runs the app
func (s *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
