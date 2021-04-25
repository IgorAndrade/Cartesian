package repository

import "cartesian/domain"

type Repository interface {
	GetAllPoints() ([]domain.Point, error)
	Insert(domain.Point) error
}
