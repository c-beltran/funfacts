package service

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/rotisserie/eris"
)

type (
	TrivialFactFinder interface {
		FindTrivialFact(ctx context.Context) (facts.Trivial, error)
	}

	TrivialFact struct {
		Finder TrivialFactFinder
	}
)

func (svc TrivialFact) Find(ctx context.Context) (facts.Trivial, error) {
	fact, err := svc.Finder.FindTrivialFact(ctx)
	if err != nil {
		return facts.Trivial{}, eris.Wrap(err, "unable to find trivial fact")
	}

	return fact, nil
}
