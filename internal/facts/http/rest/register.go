package rest

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
)

type (
	DogFacSvc interface {
		Find(ctx context.Context) (facts.Dog, error)
	}

	RegisterParams struct {
		DogFact DogFacSvc
	}
)

func Register(server *Server, services RegisterParams) {
	router := server.Router.PathPrefix("/ffact").Subrouter()
	registerDogFactHandlers(router, server, services.DogFact)
}
