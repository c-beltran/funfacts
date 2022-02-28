package rest_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/c-beltran/funfacts/internal/facts/facttesting"
	"github.com/c-beltran/funfacts/internal/facts/http/rest"
	"github.com/gorilla/mux"
)

type (
	setupServer struct {
		factSVC *facttesting.FakeFactSvc
	}
)

func newServer() (*setupServer, *rest.Server) {
	server := rest.NewServer("")

	service := &setupServer{
		factSVC: &facttesting.FakeFactSvc{},
	}

	rest.Register(server, rest.RegisterParams{
		Fact: service.factSVC,
	})

	return service, server
}

func doRequest(router *mux.Router, req *http.Request) *http.Response {
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr.Result()
}
