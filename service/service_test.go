package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/IgorAndrade/Cartesian/domain"
	"github.com/IgorAndrade/Cartesian/repository"
)

func TestPointServiceImp_GetPointsWithin(t *testing.T) {
	data := []domain.Point{{X: 67, Y: -72}, {X: 68, Y: -74}, {X: 60, Y: 72}, {X: 64, Y: -70}, {X: 65, Y: -72}}
	type fields struct {
		repo repository.Repository
	}
	type args struct {
		origin   domain.Point
		distance int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Point
		wantErr bool
	}{
		{
			name: "must return 3 points within the distance",
			fields: fields{repo: &repository.MemoryMock{
				GetFn: func() ([]domain.Point, error) {
					return data, nil
				},
			}},
			args: args{
				origin:   domain.Point{X: 66, Y: -71},
				distance: 3,
			},
			wantErr: false,
			want:    []domain.Point{{X: 67, Y: -72}, {X: 65, Y: -72}, {X: 64, Y: -70}},
		},
		{
			name: "must return 0 points within the distance",
			fields: fields{repo: &repository.MemoryMock{
				GetFn: func() ([]domain.Point, error) {
					return data, nil
				},
			}},
			args: args{
				origin:   domain.Point{X: 66, Y: 71},
				distance: 3,
			},
			wantErr: false,
			want:    []domain.Point{},
		},
		{
			name: "return error",
			fields: fields{repo: &repository.MemoryMock{
				GetFn: func() ([]domain.Point, error) {
					return nil, errors.New("any error")
				},
			}},
			args: args{
				origin:   domain.Point{X: 66, Y: 71},
				distance: 3,
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewPointService(tt.fields.repo)

			got, err := s.GetPointsWithin(tt.args.origin, tt.args.distance)
			if (err != nil) != tt.wantErr {
				t.Errorf("PointServiceImp.GetPointsWithin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointServiceImp.GetPointsWithin() = %v, want %v", got, tt.want)
			}
		})
	}
}
