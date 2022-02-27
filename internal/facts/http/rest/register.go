package rest

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o ../../facttesting/fake_fact_service.gen.go . FactSvc

type (
	// FactSvc defines the fact Service.
	FactSvc interface {
		Find(ctx context.Context, topic facts.TopicType) (facts.Topic, error)
	}

	RegisterParams struct {
		Fact FactSvc
	}
)

func Register(server *Server, services RegisterParams) {
	router := server.Router.PathPrefix("/ffact").Subrouter()
	registerDogFactHandlers(router, server, services.Fact)
	registerCatFactHandlers(router, server, services.Fact)
	registerEntertainmentFactHandlers(router, server, services.Fact)
	registerTrivialFactHandlers(router, server, services.Fact)
}
