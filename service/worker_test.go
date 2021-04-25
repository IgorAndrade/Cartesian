package service

import (
	"context"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/IgorAndrade/Cartesian/domain"
	"github.com/IgorAndrade/Cartesian/repository"
)

const jsonSample = `
[
    {
      "x": 63,
      "y": -72
    },
    {
      "x": -94,
      "y": 89
    },
    {
      "x": -30,
      "y": -38
    },
    {
      "x": -43,
      "y": -65
    },
    {
      "x": 88,
      "y": -74
    }
]
`

func TestLoader_Import(t *testing.T) {
	type fields struct {
		repo repository.Repository
	}
	type args struct {
		ctx    context.Context
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    []domain.Point
	}{
		{
			name: "test parse json",
			fields: fields{
				repo: repository.NewMemoryRepository(),
			},
			args: args{
				ctx:    context.Background(),
				reader: strings.NewReader(jsonSample),
			},
			wantErr: false,
			want:    []domain.Point{{X: 63, Y: -72}, {X: -94, Y: 89}, {X: -30, Y: -38}, {X: -43, Y: -65}, {X: 88, Y: -74}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Loader{
				repo: tt.fields.repo,
			}
			if err := l.Import(tt.args.ctx, tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("Loader.Import() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, _ := tt.fields.repo.GetAllPoints()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointServiceImp.GetPointsWithin() = %v, want %v", got, tt.want)
			}
		})
	}
}
