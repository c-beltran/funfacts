package rest_test

import (
	"github.com/c-beltran/funfacts/internal/facts/facttesting"
	"github.com/c-beltran/funfacts/internal/facts/http/rest"
)

type (
	setupServer struct {
		factSVC *facttesting.FakeFactSvc
	}
)

func newServer() (*setupServer, *rest.Server) {
	server := rest.NewServer("/")

	service := &setupServer{
		factSVC: &facttesting.FakeFactSvc{},
	}

	rest.Register(server, rest.RegisterParams{
		Fact: service.factSVC,
	})

	return service, server
}
