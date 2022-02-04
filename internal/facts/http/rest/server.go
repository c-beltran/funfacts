package rest

import "github.com/gorilla/mux"

type (
	Server struct {
		Router   *mux.Router
		BasePath string
	}
)

func NewServer(basePath string) *Server {
	return &Server{
		Router:   mux.NewRouter(),
		BasePath: basePath,
	}
}
