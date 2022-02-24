package rest

import (
	"net/http"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/gorilla/mux"
)

type (
	GetCatFactResponse struct {
		Fact string `json:"fact"`
	}
)

func registerCatFactHandlers(router *mux.Router, server *Server, svc FactSvc) {
	catRouter := router.PathPrefix("/cat").Subrouter()

	server.RegisterPath("", catRouter).Get("getCatFact", getCatFact(svc))
}

func getCatFact(svc FactSvc) func(*http.Request, *Server, map[string]interface{}) (interface{}, error) {
	return func(r *http.Request, s *Server, m map[string]interface{}) (interface{}, error) {

		fact, err := svc.Find(r.Context(), facts.TopicTypeCat)
		if err != nil {
			return nil, err
		}

		return GetCatFactResponse{
			Fact: fact.Cat,
		}, nil
	}
}
