package rest

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
)

type (
	DogFactSvc interface {
		Find(ctx context.Context) (facts.Dog, error)
	}

	CatFactSvc interface {
		Find(ctx context.Context) (facts.Cat, error)
	}

	RegisterParams struct {
		DogFact DogFactSvc
		CatFact CatFactSvc
	}
)

func Register(server *Server, services RegisterParams) {
	router := server.Router.PathPrefix("/ffact").Subrouter()
	registerDogFactHandlers(router, server, services.DogFact)
	registerCatFactHandlers(router, server, services.CatFact)
}
