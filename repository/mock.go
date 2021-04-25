package repository

import "github.com/IgorAndrade/Cartesian/domain"

type MemoryMock struct {
	GetFn    func() ([]domain.Point, error)
	InsertFn func(p domain.Point) error
}

func (m MemoryMock) GetAllPoints() ([]domain.Point, error) {
	return m.GetFn()
}
func (m *MemoryMock) Insert(p domain.Point) error {
	return m.InsertFn(p)
}
