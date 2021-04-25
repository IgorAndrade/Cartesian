package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IgorAndrade/Cartesian/repository"
	"github.com/IgorAndrade/Cartesian/server"
	"github.com/IgorAndrade/Cartesian/service"
	"golang.org/x/sync/errgroup"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	pathToLoad := os.Getenv("FILE_PATH")

	if port == "" {
		port = "8080"
	}
	if pathToLoad == "" {
		log.Fatalln("FILE_PATH is required")
	}

	ctx, cancel := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)
	defer cancel()

	r := repository.NewMemoryRepository()
	worker := service.NewWorker(r)
	if err := worker.Load(gctx, pathToLoad); err != nil {
		log.Fatalln(err)
	}

	srv := service.NewPointService(r)
	server := server.NewApi(port, srv, cancel)

	g.Go(server.Start)
	g.Go(waitSignalChannel(gctx, server.Stop))

	err := g.Wait()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Print("context was canceled")
		} else {
			fmt.Printf("received error: %v", err)
		}
	} else {
		fmt.Println("finished clean")
	}
}

func waitSignalChannel(gctx context.Context, stop func() error) func() error {
	return func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, os.Kill, syscall.SIGSEGV)
		defer stop()

		select {
		case sig := <-signalChannel:
			log.Printf("Received signal: %s\n", sig)
		case <-gctx.Done():
			log.Printf("closing signal goroutine\n")
			return gctx.Err()
		}

		return nil
	}
}
