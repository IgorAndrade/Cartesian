package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
	log.Println("Closing HTTP server")
	return a.srv.Close()
}

func (a Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Method %s Not Allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}
	x, err := validateAndGet(r, "x")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	y, err := validateAndGet(r, "y")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	distance, err := validateAndGet(r, "distance")
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

func validateAndGet(r *http.Request, param string) (int, error) {
	paramString := r.FormValue(param)
	if paramString == "" {
		return 0, errors.New("distance is required")
	}
	num, err := strconv.Atoi(paramString)
	if err != nil {
		return 0, err
	}
	return num, nil
}
