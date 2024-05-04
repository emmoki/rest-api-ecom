package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/emmoki/rest-api-ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started in %s", s.addr)

	return server.ListenAndServe()
}
