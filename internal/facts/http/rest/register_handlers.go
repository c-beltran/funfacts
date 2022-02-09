package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rotisserie/eris"
)

type (
	// GetHandler defines the HTTP GET Handlers.
	GetHandler func(r *http.Request, s *Server, args map[string]interface{}) (interface{}, error)

	// PostHandler definres the HTTP POST Handlers.
	PostHandler func(r *http.Request, s *Server, args map[string]interface{}) (string, error)

	respondFunc func(w http.ResponseWriter, payload interface{}, err error)
)

func (s *Server) RegisterGetHandler(path string, router *mux.Router, handler GetHandler, args []string) *mux.Route {
	return s.registerGet(path, router, handler, args, responseStatus)
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

func responseStatus(w http.ResponseWriter, payload interface{}, err error) {
	if payload == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		responseCode := http.StatusOK

		res, err := json.Marshal(payload)
		if err != nil {
			errToJson := eris.ToJSON(err, true)
			res, _ = json.Marshal(errToJson)
			responseCode = http.StatusInternalServerError
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseCode)

		if _, err := w.Write(res); err != nil {
			fmt.Fprintf(os.Stderr, "error %s", err)
		}
	}
}

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
