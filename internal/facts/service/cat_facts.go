package service

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/rotisserie/eris"
)

type (
	CatFactFinder interface {
		FindCatFact(ctx context.Context) (facts.Diversity, error)
	}

	CatFact struct {
		Finder CatFactFinder
	}
)

func (svc CatFact) Find(ctx context.Context) (facts.Diversity, error) {
	fact, err := svc.Finder.FindCatFact(ctx)
	if err != nil {
		return facts.Diversity{}, eris.Wrap(err, "unable to find cat fact")
	}

	return fact, nil
}
