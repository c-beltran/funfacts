package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	GetEntertainmentFactResponse struct {
		Fact string `json:"fact"`
	}
)

func registerEntertainmentFactHandlers(router *mux.Router, server *Server, svc EntertainmentFactSvc) {
	entRouter := router.PathPrefix("/entertainment").Subrouter()

	server.RegisterPath("", entRouter).Get("getEntertainmentFact", getEntertainmentFact(svc))
}

func getEntertainmentFact(svc EntertainmentFactSvc) func(*http.Request, *Server, map[string]interface{}) (interface{}, error) {
	return func(r *http.Request, s *Server, m map[string]interface{}) (interface{}, error) {

		fact, err := svc.Find(r.Context())
		if err != nil {
			return nil, err
		}

		return GetEntertainmentFactResponse{
			Fact: fact.Entertainment,
		}, nil
	}
}
