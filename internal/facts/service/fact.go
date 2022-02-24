package service

import (
	"context"

	"github.com/c-beltran/funfacts/internal/facts"
	"github.com/rotisserie/eris"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o ../facttesting/fake_cat_fact_finder.gen.go . CatFactFinder
//counterfeiter:generate -o ../facttesting/fake_dog_fact_finder.gen.go . DogFactFinder
//counterfeiter:generate -o ../facttesting/fake_entertainment_fact_finder.gen.go . EntertainmentFactFinder
//counterfeiter:generate -o ../facttesting/fake_trivial_fact_finder.gen.go . TrivialFactFinder

type (
	CatFactFinder interface {
		FindCatFact(ctx context.Context) (facts.Topic, error)
	}

	DogFactFinder interface {
		FindDogFact(ctx context.Context) (facts.Topic, error)
	}

	EntertainmentFactFinder interface {
		FindEntertainmentFact(ctx context.Context) (facts.Topic, error)
	}

	TrivialFactFinder interface {
		FindTrivialFact(ctx context.Context) (facts.Topic, error)
	}

	FactSVC struct {
		CFinder CatFactFinder
		DFinder DogFactFinder
		EFinder EntertainmentFactFinder
		TFinder TrivialFactFinder
	}
)

func (svc FactSVC) Find(ctx context.Context, topic string) (facts.Topic, error) {
	var (
		fact facts.Topic
		err  error
	)

	switch topic {
	case "cat":
		fact, err = svc.CFinder.FindCatFact(ctx)
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "unable to find cat fact")
		}
	case "dog":
		fact, err = svc.DFinder.FindDogFact(ctx)
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "unable to find dog fact")
		}
	case "entertainment":
		fact, err = svc.EFinder.FindEntertainmentFact(ctx)
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "unable to find entertainment fact")
		}
	case "trivial":
		fact, err = svc.TFinder.FindTrivialFact(ctx)
		if err != nil {
			return facts.Topic{}, eris.Wrap(err, "unable to find trivial fact")
		}
	default:
		return facts.Topic{}, eris.New("unknown topic")
	}

	return fact, nil
}
