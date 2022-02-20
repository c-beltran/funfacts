package service

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/rotisserie/eris"
)

type (
	EntertainmentFactFinder interface {
		FindEntertainmentFact(ctx context.Context) (facts.Entertainment, error)
	}

	EntertainmentFact struct {
		Finder EntertainmentFactFinder
	}
)

func (svc EntertainmentFact) Find(ctx context.Context) (facts.Entertainment, error) {
	fact, err := svc.Finder.FindEntertainmentFact(ctx)
	if err != nil {
		return facts.Entertainment{}, eris.Wrap(err, "unable to find entertainment fact")
	}

	return fact, nil
}