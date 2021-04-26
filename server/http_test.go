package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorAndrade/Cartesian/domain"
	"github.com/IgorAndrade/Cartesian/service"
)

func TestApi_ServeHTTP(t *testing.T) {
	type fields struct {
		service service.PointService
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedStatus int
	}{
		{
			name: "expected Status 200",
			fields: fields{
				service: service.PointServiceMock{
					GetFn: func(p domain.Point, i int) ([]domain.Point, error) {
						return []domain.Point{}, nil
					},
				},
			},
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/api/points?x=65&y=-70&distance=50", nil),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "expected Status 400, missing y param",
			fields: fields{
				service: service.PointServiceMock{
					GetFn: func(p domain.Point, i int) ([]domain.Point, error) {
						return []domain.Point{}, nil
					},
				},
			},
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/api/points?x=65&distance=50", nil),
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "expected Status 400, missing x param",
			fields: fields{
				service: service.PointServiceMock{
					GetFn: func(p domain.Point, i int) ([]domain.Point, error) {
						return []domain.Point{}, nil
					},
				},
			},
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/api/points?y=65&distance=50", nil),
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "expected Status 400, missing distance param",
			fields: fields{
				service: service.PointServiceMock{
					GetFn: func(p domain.Point, i int) ([]domain.Point, error) {
						return []domain.Point{}, nil
					},
				},
			},
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/api/points?x=65&y=65", nil),
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "expected Status 400, distance not num",
			fields: fields{
				service: service.PointServiceMock{
					GetFn: func(p domain.Point, i int) ([]domain.Point, error) {
						return []domain.Point{}, nil
					},
				},
			},
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/api/points?x=65&y=65&&distance=50m", nil),
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "expected Status 400, Method Not Allowed",
			fields: fields{
				service: service.PointServiceMock{
					GetFn: func(p domain.Point, i int) ([]domain.Point, error) {
						return []domain.Point{}, nil
					},
				},
			},
			args: args{
				r: httptest.NewRequest(http.MethodPut, "/api/points?x=65&distance=50", nil),
			},
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name: "expected Status 500, service return error",
			fields: fields{
				service: service.PointServiceMock{
					GetFn: func(p domain.Point, i int) ([]domain.Point, error) {
						return []domain.Point{}, errors.New("any error")
					},
				},
			},
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/api/points?x=65&y=65&distance=50", nil),
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewApi("8080", tt.fields.service, nil)
			rec := httptest.NewRecorder()
			a.ServeHTTP(rec, tt.args.r)
			if tt.expectedStatus != rec.Code {
				t.Errorf("[ServeHTTP] expectedStatus = %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}
