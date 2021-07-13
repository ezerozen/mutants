package mutants

import (
	"context"
	"erozen/mutants/internal/db"
)

type Repository interface {
	Save(ctx context.Context, test *db.Test) error
	Count(ctx context.Context, isMutant bool) (int64, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) IsMutant(ctx context.Context, dna []string) (bool, error) {
	test := db.Test{
		DNA:    dna,
		Mutant: isMutant(dna),
	}

	if err := s.repository.Save(ctx, &test); err != nil {
		return false, err
	}

	return test.Mutant, nil
}

func (s *Service) Stats(ctx context.Context) (Stats, error) {
	mutantCount, err := s.repository.Count(ctx, true)
	if err != nil {
		return Stats{}, err
	}

	humanCount, err := s.repository.Count(ctx, false)
	if err != nil {
		return Stats{}, err
	}

	return Stats{
		MutantDNACount: mutantCount,
		HumanDNACount:  humanCount,
		Ratio:          float64(mutantCount / humanCount),
	}, nil
}
