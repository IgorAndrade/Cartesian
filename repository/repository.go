package repository

import "github.com/IgorAndrade/Cartesian/domain"

type Repository interface {
	GetAllPoints() ([]domain.Point, error)
	Insert(domain.Point) error
}

type Memory struct {
	data []domain.Point
}

func (m Memory) GetAllPoints() ([]domain.Point, error) {
	r := make([]domain.Point, len(m.data))
	copy(r, m.data)
	return r, nil
}
func (m *Memory) Insert(p domain.Point) error {
	m.data = append(m.data, p)
	return nil
}

func NewMemoryRepository() Repository {
	return &Memory{}
}
