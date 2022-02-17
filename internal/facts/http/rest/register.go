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

	TrivialFactSvc interface {
		Find(ctx context.Context) (facts.Trivial, error)
	}

	RegisterParams struct {
		DogFact           DogFactSvc
		CatFact           CatFactSvc
		EntertainmentFact EntertainmentFactSvc
		TrivialFact       TrivialFactSvc
	}
)

func Register(server *Server, services RegisterParams) {
	router := server.Router.PathPrefix("/ffact").Subrouter()
	registerDogFactHandlers(router, server, services.DogFact)
	registerCatFactHandlers(router, server, services.CatFact)
	registerEntertainmentFactHandlers(router, server, services.EntertainmentFact)
	registerTrivialFactHandlers(router, server, services.TrivialFact)
}
