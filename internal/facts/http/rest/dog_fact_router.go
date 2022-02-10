package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	GetDogFactResponse struct {
		Fact string `json:"fact"`
	}
)

func registerDogFactHandlers(router *mux.Router, server *Server, svc DogFacSvc) {
	dogRouter := router.PathPrefix("/dog").Subrouter()

	server.RegisterPath("", dogRouter).Get("getDogFact", getDogFact(svc))
}

func getDogFact(svc DogFacSvc) func(*http.Request, *Server, map[string]interface{}) (interface{}, error) {
	return func(r *http.Request, s *Server, m map[string]interface{}) (interface{}, error) {

		dog, err := svc.Find(r.Context())
		if err != nil {
			return nil, err
		}

		return GetDogFactResponse{
			Fact: dog.Fact,
		}, nil
	}
}