package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IgorAndrade/Cartesian/domain"
	"github.com/IgorAndrade/Cartesian/service"
)

type Api struct {
	service service.PointService
	srv     *http.Server
	cancel  context.CancelFunc
}

func (a Api) Start() error {
	defer a.cancel()
	return a.srv.ListenAndServe()
}

func (a Api) Stop() error {
	return a.srv.Close()
}

func (a Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paramX := r.FormValue("x")
	if paramX == "" {
		http.Error(w, "x is required", http.StatusBadRequest)
		return
	}
	paramY := r.FormValue("y")
	if paramY == "" {
		http.Error(w, "y is required", http.StatusBadRequest)
		return
	}
	paramDistance := r.FormValue("distance")
	if paramDistance == "" {
		http.Error(w, "distance is required", http.StatusBadRequest)
		return
	}
	x, err := strconv.Atoi(paramX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(paramY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	distance, err := strconv.Atoi(paramDistance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err := a.service.GetPointsWithin(domain.Point{X: x, Y: y}, distance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)

}

func NewApi(port string, service service.PointService, cancel context.CancelFunc) *Api {
	api := &Api{service: service, cancel: cancel}
	mux := http.NewServeMux()
	mux.Handle("/api/points", http.HandlerFunc(api.ServeHTTP))

	srv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: mux}
	api.srv = srv
	return api
}
