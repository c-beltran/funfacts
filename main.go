package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/c-beltran/funfacts/internal/facts/apis"
	"github.com/c-beltran/funfacts/internal/facts/http/rest"
	"github.com/c-beltran/funfacts/internal/facts/service"
	"github.com/gorilla/mux"
)

func main() {
	httpClient := http.Client{
		Timeout:   http.DefaultClient.Timeout,
		Transport: http.DefaultTransport,
	}

	dogAPI := apis.NewClient(&httpClient, "https://dog-api.kinduff.com")

	//-
	restServer := rest.NewServer("/")
	rest.Register(restServer, rest.RegisterParams{
		DogFact: &service.DogFact{
			Finder: dogAPI,
		},
	})

	err := restServer.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe("0.0.0.0:8080", restServer.Router)
}
