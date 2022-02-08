package rest

import "github.com/gorilla/mux"

type (
	// RegisteredPath defines the type used when a path is registered.
	RegisteredPath struct {
		args   []string
		path   string
		router *mux.Router
		server *Server
	}
)

func (s *Server) RegisterPath(path string, router *mux.Router, args ...string) *RegisteredPath {
	return &RegisteredPath{
		args:   args,
		path:   path,
		router: router,
		server: s,
	}
}

func (r *RegisteredPath) Get(opID string, handler GetHandler) *RegisteredPath {
	r.server.RegisterGetHandler(r.path, r.router, handler, r.args).Name(opID)
	return r
}

// func (r *RegisteredPath) Post(opID string, handler PostHandler) *RegisteredPath {
// 	r.server.RegisterPostHandler(r.path, r.router, handler, r.args).Name(opID)
// }
