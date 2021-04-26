package service

import "github.com/IgorAndrade/Cartesian/domain"

type PointServiceMock struct {
	GetFn func(domain.Point, int) ([]domain.Point, error)
}

func (p PointServiceMock) GetPointsWithin(point domain.Point, distance int) ([]domain.Point, error) {
	return p.GetFn(point, distance)
}
