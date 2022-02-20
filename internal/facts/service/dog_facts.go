package service

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/rotisserie/eris"
)

type (
	DogFactFinder interface {
		FindDogFact(ctx context.Context) (facts.Diversity, error)
	}

	DogFact struct {
		Finder DogFactFinder
	}
)

func (svc DogFact) Find(ctx context.Context) (facts.Diversity, error) {
	fact, err := svc.Finder.FindDogFact(ctx)
	if err != nil {
		return facts.Diversity{}, eris.Wrap(err, "unable to find dog fact")
	}

	return fact, nil
}
