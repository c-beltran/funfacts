package rest

import (
	"net/http"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/gorilla/mux"
)

type (
	GetTrivialFactResponse struct {
		Fact string `json:"fact"`
	}
)

func registerTrivialFactHandlers(router *mux.Router, server *Server, svc FactSvc) {
	trivRouter := router.PathPrefix("/trivial").Subrouter()

	server.RegisterPath("", trivRouter).Get("getTrivialFact", getTrivialFact(svc))
}

func getTrivialFact(svc FactSvc) func(*http.Request, *Server, map[string]interface{}) (interface{}, error) {
	return func(r *http.Request, s *Server, m map[string]interface{}) (interface{}, error) {

		fact, err := svc.Find(r.Context(), facts.TopicTypeTrivial)
		if err != nil {
			return nil, err
		}

		return GetTrivialFactResponse{
			Fact: fact.Trivial,
		}, nil
	}
}
