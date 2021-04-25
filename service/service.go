package service

import (
	"sort"

	"github.com/IgorAndrade/Cartesian/domain"
	"github.com/IgorAndrade/Cartesian/repository"
)

type PointService interface {
	GetPointsWithin(domain.Point, int) ([]domain.Point, error)
}

type PointServiceImp struct {
	repo repository.Repository
}

func (s PointServiceImp) GetPointsWithin(origin domain.Point, distance int) ([]domain.Point, error) {
	listP, err := s.repo.GetAllPoints()
	if err != nil {
		return nil, err
	}

	r := make([]domain.Point, 0)
	for _, p := range listP {
		if p.DistanceTo(origin) <= distance {
			r = append(r, p)
		}
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i].DistanceTo(origin) < r[j].DistanceTo(origin)
	})

	return r, nil
}

func NewPointService(repo repository.Repository) PointService {
	return &PointServiceImp{repo: repo}
}
