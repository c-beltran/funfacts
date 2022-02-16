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

	EntertainmentFactSvc interface {
		Find(ctx context.Context) (facts.Entertainment, error)
	}

	RegisterParams struct {
		DogFact           DogFactSvc
		CatFact           CatFactSvc
		EntertainmentFact EntertainmentFactSvc
	}
)

func Register(server *Server, services RegisterParams) {
	router := server.Router.PathPrefix("/ffact").Subrouter()
	registerDogFactHandlers(router, server, services.DogFact)
	registerCatFactHandlers(router, server, services.CatFact)
	registerEntertainmentFactHandlers(router, server, services.EntertainmentFact)
}
