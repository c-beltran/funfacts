package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/c-beltran/funfacts/internal/facts/http/rest"
	"github.com/google/go-cmp/cmp"
)

func Test_getCartAsset(t *testing.T) {
	t.Parallel()

	type (
		// svcArgs struct {
		// 	ID int64
		// }

		responseUnion struct {
			Error string `json:"error"`
			rest.GetCatFactResponse
		}

		output struct {
			statusCode int
			response   responseUnion
			// serviceArgs *svcArgs
		}

		input struct {
			svc func(*setupServer)
		}
	)

	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			"OK",
			input{
				svc: func(s *setupServer) {
					s.factSVC.FindReturns(facts.Topic{
						Cat: "meow",
					}, nil)
				},
			},
			output{
				statusCode: http.StatusOK,
				response: responseUnion{
					GetCatFactResponse: rest.GetCatFactResponse{
						Fact: "meow",
					},
				},
				// serviceArgs: &svcArgs{
				// 	ID: 10,
				// },
			},
		},
		{
			"Service error",
			input{
				svc: func(s *setupServer) {
					s.factSVC.FindReturns(facts.Topic{}, errors.New("FactSvc.CatFact error"))
				},
			},
			output{
				statusCode: http.StatusInternalServerError,
				response: responseUnion{
					Error: "FactSvc.CatFact error",
				},
				// serviceArgs: &svcArgs{
				// 	ID: 10,
				// },
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s, server := newServer()
			tt.input.svc(s)

			res := httptest.NewRequest(server.Router, t, http.MethodGet, "/delivery/assets/carts/10")

			resttesting.TestResponseCode(t, tt.output.statusCode, res.StatusCode)
			resttesting.TestResponse(t, res.Body, &tt.output.response, new(responseUnion))

			if tt.output.serviceArgs != nil {
				_, ID := s.cartAsset.CartArgsForCall(0)
				if diff := cmp.Diff(*tt.output.serviceArgs, svcArgs{ID}); diff != "" {
					t.Errorf("CartArgsForCall don't match:\n%s", diff)
				}
			}
		})
	}
}
