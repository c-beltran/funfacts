package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	GetTrivialFactResponse struct {
		Fact string `json:"fact"`
	}
)

func registerTrivialFactHandlers(router *mux.Router, server *Server, svc TrivialFactSvc) {
	catRouter := router.PathPrefix("/trivial").Subrouter()

	server.RegisterPath("", catRouter).Get("getTrivialFact", getTrivialFact(svc))
}

func getTrivialFact(svc TrivialFactSvc) func(*http.Request, *Server, map[string]interface{}) (interface{}, error) {
	return func(r *http.Request, s *Server, m map[string]interface{}) (interface{}, error) {

		cat, err := svc.Find(r.Context())
		if err != nil {
			return nil, err
		}

		return GetTrivialFactResponse{
			Fact: cat.Fact,
		}, nil
	}
}
