package apis_test

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/c-beltran/funfacts/internal/facts/apis"
	"github.com/google/go-cmp/cmp"
	"github.com/rotisserie/eris"
)

func Test_FindDogFact(t *testing.T) {
	type (
		setup struct {
			host       string
			statusCode int
			response   string
			ctx        context.Context
		}

		expected struct {
			fact facts.Topic
			err  error
		}
	)

	tests := []struct {
		name     string
		setup    setup
		expected expected
	}{
		{
			"ok",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{
					"facts": [
						"Dogs wag their tail when they are happy"
					],
					"success": true
				}`,
				context.Background(),
			},
			expected{
				fact: facts.Topic{
					Dog: "Dogs wag their tail when they are happy",
				},
			},
		},
		{
			"bad host",
			setup{
				ctx: context.Background(),
			},
			expected{
				err: eris.New(`unsupported protocol scheme ""`),
			},
		},
		{
			"bad status",
			setup{
				host:       "http://ffun.com",
				statusCode: http.StatusInternalServerError,
				ctx:        context.Background(),
			},
			expected{
				err: eris.New("bad status code from server 500"),
			},
		},
		{
			"bad request",
			setup{
				host: "http://ffun.com",
			},
			expected{
				err: eris.New(`net/http: nil Context`),
			},
		},
		{
			"invalid respose",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{\\,m`,
				context.Background(),
			},
			expected{
				err: eris.New(`invalid character '\\' looking for beginning of object key string`),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.WriteHeader(test.setup.statusCode)
				w.Write([]byte(test.setup.response))
			})

			s := httptest.NewServer(h)
			defer s.Close()

			c := &http.Client{
				Transport: &http.Transport{
					DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
						return net.Dial(network, s.Listener.Addr().String())
					},
				},
			}

			api := apis.NewClient(c, test.setup.host)

			actual, err := api.FindDogFact(test.setup.ctx)
			if !eris.Is(test.expected.err, eris.Cause(err)) {
				t.Fatalf("expected %v, got %v", test.expected.err, eris.Cause(err))
			}

			if diff := cmp.Diff(test.expected.fact, actual); diff != "" {
				t.Error("expected result does not match\n", diff)
			}
		})
	}
}

func Test_FindCatFact(t *testing.T) {
	type (
		setup struct {
			host       string
			statusCode int
			response   string
			ctx        context.Context
		}

		expected struct {
			fact facts.Topic
			err  error
		}
	)

	tests := []struct {
		name     string
		setup    setup
		expected expected
	}{
		{
			"ok",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{
					"data": [
						"Cats are felines"
					]
				}`,
				context.Background(),
			},
			expected{
				fact: facts.Topic{
					Cat: "Cats are felines",
				},
			},
		},
		{
			"bad host",
			setup{
				ctx: context.Background(),
			},
			expected{
				err: eris.New(`unsupported protocol scheme ""`),
			},
		},
		{
			"bad status",
			setup{
				host:       "http://ffun.com",
				statusCode: http.StatusInternalServerError,
				ctx:        context.Background(),
			},
			expected{
				err: eris.New("bad status code from server 500"),
			},
		},
		{
			"bad request",
			setup{
				host: "http://ffun.com",
			},
			expected{
				err: eris.New(`net/http: nil Context`),
			},
		},
		{
			"invalid respose",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{\\,m`,
				context.Background(),
			},
			expected{
				err: eris.New(`invalid character '\\' looking for beginning of object key string`),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.WriteHeader(test.setup.statusCode)
				w.Write([]byte(test.setup.response))
			})

			s := httptest.NewServer(h)
			defer s.Close()

			c := &http.Client{
				Transport: &http.Transport{
					DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
						return net.Dial(network, s.Listener.Addr().String())
					},
				},
			}

			api := apis.NewClient(c, test.setup.host)

			actual, err := api.FindCatFact(test.setup.ctx)
			if !eris.Is(test.expected.err, eris.Cause(err)) {
				t.Fatalf("expected %v, got %v", test.expected.err, eris.Cause(err))
			}

			if diff := cmp.Diff(test.expected.fact, actual); diff != "" {
				t.Error("expected result does not match\n", diff)
			}
		})
	}
}

func Test_FindEntertainmentFact(t *testing.T) {
	type (
		setup struct {
			host       string
			statusCode int
			response   string
			ctx        context.Context
		}

		expected struct {
			fact facts.Topic
			err  error
		}
	)

	tests := []struct {
		name     string
		setup    setup
		expected expected
	}{
		{
			"ok",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{
					"data": {
						"id": "100",
						"fact": "Water is made out of 2 hydrogen molecules and one oxygen molecule",
						"category": "science"
					}
				}`,
				context.Background(),
			},
			expected{
				fact: facts.Topic{
					Entertainment: "Water is made out of 2 hydrogen molecules and one oxygen molecule",
				},
			},
		},
		{
			"bad host",
			setup{
				ctx: context.Background(),
			},
			expected{
				err: eris.New(`unsupported protocol scheme ""`),
			},
		},
		{
			"bad status",
			setup{
				host:       "http://ffun.com",
				statusCode: http.StatusInternalServerError,
				ctx:        context.Background(),
			},
			expected{
				err: eris.New("bad status code from server 500"),
			},
		},
		{
			"bad request",
			setup{
				host: "http://ffun.com",
			},
			expected{
				err: eris.New(`net/http: nil Context`),
			},
		},
		{
			"invalid respose",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{\\,m`,
				context.Background(),
			},
			expected{
				err: eris.New(`invalid character '\\' looking for beginning of object key string`),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.WriteHeader(test.setup.statusCode)
				w.Write([]byte(test.setup.response))
			})

			s := httptest.NewServer(h)
			defer s.Close()

			c := &http.Client{
				Transport: &http.Transport{
					DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
						return net.Dial(network, s.Listener.Addr().String())
					},
				},
			}

			api := apis.NewClient(c, test.setup.host)

			actual, err := api.FindEntertainmentFact(test.setup.ctx)
			if !eris.Is(test.expected.err, eris.Cause(err)) {
				t.Fatalf("expected %v, got %v", test.expected.err, eris.Cause(err))
			}

			if diff := cmp.Diff(test.expected.fact, actual); diff != "" {
				t.Error("expected result does not match\n", diff)
			}
		})
	}
}

func Test_FindTrivialFact(t *testing.T) {
	type (
		setup struct {
			host       string
			statusCode int
			response   string
			ctx        context.Context
		}

		expected struct {
			fact facts.Topic
			err  error
		}
	)

	tests := []struct {
		name     string
		setup    setup
		expected expected
	}{
		{
			"ok",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{
					"text": "On every continent there is a city called Rome"
				}`,
				context.Background(),
			},
			expected{
				fact: facts.Topic{
					Trivial: "On every continent there is a city called Rome",
				},
			},
		},
		{
			"bad host",
			setup{
				ctx: context.Background(),
			},
			expected{
				err: eris.New(`unsupported protocol scheme ""`),
			},
		},
		{
			"bad status",
			setup{
				host:       "http://ffun.com",
				statusCode: http.StatusInternalServerError,
				ctx:        context.Background(),
			},
			expected{
				err: eris.New("bad status code from server 500"),
			},
		},
		{
			"bad request",
			setup{
				host: "http://ffun.com",
			},
			expected{
				err: eris.New(`net/http: nil Context`),
			},
		},
		{
			"invalid respose",
			setup{
				"http://ffun.com",
				http.StatusOK,
				`{\\,m`,
				context.Background(),
			},
			expected{
				err: eris.New(`invalid character '\\' looking for beginning of object key string`),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.WriteHeader(test.setup.statusCode)
				w.Write([]byte(test.setup.response))
			})

			s := httptest.NewServer(h)
			defer s.Close()

			c := &http.Client{
				Transport: &http.Transport{
					DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
						return net.Dial(network, s.Listener.Addr().String())
					},
				},
			}

			api := apis.NewClient(c, test.setup.host)

			actual, err := api.FindTrivialFact(test.setup.ctx)
			if !eris.Is(test.expected.err, eris.Cause(err)) {
				t.Fatalf("expected %v, got %v", test.expected.err, eris.Cause(err))
			}

			if diff := cmp.Diff(test.expected.fact, actual); diff != "" {
				t.Error("expected result does not match\n", diff)
			}
		})
	}
}
