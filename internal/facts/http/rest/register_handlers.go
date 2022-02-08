package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rotisserie/eris"
)

type (
	// GetHandler defines the HTTP GET Handlers.
	GetHandler func(r *http.Request, s *Server, args map[string]interface{}) (interface{}, error)

	// PostHandler definres the HTTP POST Handlers.
	PostHandler func(r *http.Request, s *Server, args map[string]interface{}) (string, error)

	respondFunc func(w http.ResponseWriter, payload interface{}, err error) (string, error)
)

func (s *Server) RegisterGetHandler(path string, router *mux.Router, handler GetHandler, args ...string) *mux.Route {
	return s.registerGet()
}

func (s *Server) registerGet(path string, router *mux.Router, handler GetHandler, args []string, respond respondFunc) *mux.Route {
	return router.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		args, err := varParams(r, args...)
		if err != nil {
			log.Print(err)
			return
		}

		result, err := handler(r, s, args)

		respond(rw, result, err)
	}).Methods(http.MethodGet)
}

//-

func varParams(r *http.Request, names ...string) (map[string]interface{}, error) {
	result := make(map[string]interface{}, len(names))
	routeVars := mux.Vars(r)

	for n := range names {
		str, valid := routeVars[names[n]]
		if !valid {
			return nil, eris.New("variable params not found")
		}

		val, err := strconv.Atoi(str)
		if err != nil {
			break
		}
		result[names[n]] = val
	}

	return result, nil
}
