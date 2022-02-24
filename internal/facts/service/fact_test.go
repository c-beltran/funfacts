package service_test

import (
	"context"
	"testing"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/c-beltran/funfacts/internal/facts/facttesting"
	"github.com/c-beltran/funfacts/internal/facts/service"
	"github.com/google/go-cmp/cmp"
	"github.com/rotisserie/eris"
)

type setupFactSvc struct {
	CFinder *facttesting.FakeCatFactFinder
	DFinder *facttesting.FakeDogFactFinder
	EFinder *facttesting.FakeEntertainmentFactFinder
	TFinder *facttesting.FakeTrivialFactFinder
}

func TestFactSvc_Find(t *testing.T) {
	type (
		expected struct {
			fact facts.Topic
			err  error
		}

		input struct {
			topic string
		}
	)

	tests := []struct {
		name     string
		setup    func(*setupFactSvc)
		input    input
		expected expected
	}{
		{
			name: "OK: Dog fact",
			setup: func(s *setupFactSvc) {
				s.DFinder.FindDogFactReturns(facts.Topic{
					Dog: "Some dog fact",
				}, nil)
			},
			input: input{
				topic: "dog",
			},
			expected: expected{
				fact: facts.Topic{
					Dog: "Some dog fact",
				},
				err: nil,
			},
		},
		{
			name: "OK: Cat fact",
			setup: func(s *setupFactSvc) {
				s.CFinder.FindCatFactReturns(facts.Topic{
					Cat: "Some cat fact",
				}, nil)
			},
			input: input{
				topic: "cat",
			},
			expected: expected{
				fact: facts.Topic{
					Cat: "Some cat fact",
				},
				err: nil,
			},
		},
		{
			name: "OK: Entertainment fact",
			setup: func(s *setupFactSvc) {
				s.EFinder.FindEntertainmentFactReturns(facts.Topic{
					Entertainment: "Some entertainment fact",
				}, nil)
			},
			input: input{
				topic: "entertainment",
			},
			expected: expected{
				fact: facts.Topic{
					Entertainment: "Some entertainment fact",
				},
				err: nil,
			},
		},
		{
			name: "OK: Trivial fact",
			setup: func(s *setupFactSvc) {
				s.TFinder.FindTrivialFactReturns(facts.Topic{
					Trivial: "Some trivial fact",
				}, nil)
			},
			input: input{
				topic: "trivial",
			},
			expected: expected{
				fact: facts.Topic{
					Trivial: "Some trivial fact",
				},
				err: nil,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			s, svc := newFactSvc(t)
			test.setup(s)

			got, err := svc.Find(context.Background(), test.input.topic)
			if !eris.Is(test.expected.err, eris.Cause(err)) {
				t.Fatalf("expected %v, got %v", test.expected.err, eris.Cause(err))
			}

			if diff := cmp.Diff(test.expected.fact, got); diff != "" {
				t.Error("expected result does not match\n", diff)
			}
		})
	}
}

func newFactSvc(t *testing.T) (*setupFactSvc, *service.FactSVC) {
	s := &setupFactSvc{
		CFinder: &facttesting.FakeCatFactFinder{},
		DFinder: &facttesting.FakeDogFactFinder{},
		EFinder: &facttesting.FakeEntertainmentFactFinder{},
		TFinder: &facttesting.FakeTrivialFactFinder{},
	}

	svc := &service.FactSVC{
		DFinder: s.DFinder,
		CFinder: s.CFinder,
		EFinder: s.EFinder,
		TFinder: s.TFinder,
	}

	return s, svc
}
