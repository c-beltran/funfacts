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

func Test_getTrivialAsset(t *testing.T) {
	t.Parallel()

	type (
		responseUnion struct {
			Error string `json:"error"`
			rest.GetTrivialFactResponse
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
						Trivial: "random",
					}, nil)
				},
			},
			output{
				statusCode: http.StatusOK,
				response: responseUnion{
					GetTrivialFactResponse: rest.GetTrivialFactResponse{
						Fact: "random",
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

			res := doRequest(server.Router, httptest.NewRequest(http.MethodGet, "/ffact/trivial", nil))

			trivialFact := rest.GetTrivialFactResponse{}

			if err := json.NewDecoder(res.Body).Decode(&trivialFact); err != nil {
				t.Fatalf("error unmarshaling body: %s", err)
			}
			defer res.Body.Close()

			if res.StatusCode != tt.output.statusCode {
				t.Errorf("response status does not match: %d | %d", res.StatusCode, tt.output.statusCode)
			}

			if diff := cmp.Diff(tt.output.response.GetTrivialFactResponse, trivialFact); diff != "" {
				t.Errorf("response doesn't match:\n%s", diff)
			}
		})
	}
}
