package rest_test

import (
	"encoding/json"
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
		responseUnion struct {
			Error string `json:"error"`
			rest.GetCatFactResponse
		}

		output struct {
			statusCode int
			response   responseUnion
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
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s, server := newServer()
			tt.input.svc(s)

			res := doRequest(server.Router, httptest.NewRequest(http.MethodGet, "/ffact/cat", nil))

			catFact := rest.GetCatFactResponse{}

			if err := json.NewDecoder(res.Body).Decode(&catFact); err != nil {
				t.Fatalf("error unmarshaling body: %s", err)
			}
			defer res.Body.Close()

			if res.StatusCode != tt.output.statusCode {
				t.Errorf("response status does not match: %d | %d", res.StatusCode, tt.output.statusCode)
			}

			if diff := cmp.Diff(tt.output.response.GetCatFactResponse, catFact); diff != "" {
				t.Errorf("response doesn't match:\n%s", diff)
			}
		})
	}
}
