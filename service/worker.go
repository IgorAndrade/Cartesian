package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/IgorAndrade/Cartesian/domain"
	"github.com/IgorAndrade/Cartesian/repository"
)

type Loader struct {
	repo repository.Repository
}

func (l Loader) Load(ctx context.Context, path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error when open file %s \n %v\n", path, err)
		return err
	}

	return l.Import(ctx, file)
}

func (l Loader) Import(ctx context.Context, reader io.Reader) error {
	dec := json.NewDecoder(reader)
	if _, err := dec.Token(); err != nil {
		log.Println("Error parsing token: ", err)
		return err
	}
	count := 0
	for dec.More() {
		select {
		case <-ctx.Done():
			log.Println("Sending canceled")
			return nil

		default:
			var point domain.Point
			if err := dec.Decode(&point); err != nil {
				log.Println("Error parsing point: ", err)
				return err
			}
			if err := l.repo.Insert(point); err != nil {
				return err
			}
			count++
		}
	}
	log.Printf("Imported %d points", count)

	return nil
}

func NewWorker(r repository.Repository) *Loader {
	return &Loader{repo: r}
}
