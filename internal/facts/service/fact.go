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
		FindCatFact(ctx context.Context) (facts.FactTopic, error)
	}

	DogFactFinder interface {
		FindDogFact(ctx context.Context) (facts.FactTopic, error)
	}

	EntertainmentFactFinder interface {
		FindEntertainmentFact(ctx context.Context) (facts.FactTopic, error)
	}

	TrivialFactFinder interface {
		FindTrivialFact(ctx context.Context) (facts.FactTopic, error)
	}

	FactSVC struct {
		DFinder DogFactFinder
		CFinder CatFactFinder
		EFinder EntertainmentFactFinder
		TFinder TrivialFactFinder
	}
)

func (svc FactSVC) Find(ctx context.Context, topic string) (facts.FactTopic, error) {
	var (
		fact facts.FactTopic
		err  error
	)

	switch topic {
	case "cat":
		fact, err = svc.CFinder.FindCatFact(ctx)
		if err != nil {
			return facts.FactTopic{}, eris.Wrap(err, "unable to find cat fact")
		}
	case "dog":
		fact, err = svc.DFinder.FindDogFact(ctx)
		if err != nil {
			return facts.FactTopic{}, eris.Wrap(err, "unable to find dog fact")
		}
	case "entertainment":
		fact, err = svc.EFinder.FindEntertainmentFact(ctx)
		if err != nil {
			return facts.FactTopic{}, eris.Wrap(err, "unable to find entertainment fact")
		}
	case "trivial":
		fact, err = svc.TFinder.FindTrivialFact(ctx)
		if err != nil {
			return facts.FactTopic{}, eris.Wrap(err, "unable to find trivial fact")
		}
	}

	return fact, nil
}
