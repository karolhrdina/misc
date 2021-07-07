package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/karolhrdina/misc/hw/pkg/env"
	"github.com/karolhrdina/misc/hw/services/client-api/service"

	"github.com/braintree/manners"
	"github.com/pkg/errors"
)

var (
	apiAddr     = env.GetVar("API_ADDR", "0.0.0.0:8086")
	readTimeout = 60 * time.Second
)

func main() {
	if err := runHTTP(); err != nil {
		log.Fatalf("run error: %s", err.Error())
	}
}

func runHTTP() error {
	service, err := service.InitializeService()
	if err != nil {
		return errors.Wrap(err, "unable to initialize service")
	}

	defer func() {
		errC := service.Close()
		if errC != nil {
			log.Printf("error closing service: %s", errC.Error())
		}
	}()

	apiService := manners.NewWithServer(&http.Server{
		Addr:        apiAddr,
		ReadTimeout: readTimeout,
		Handler:     service.Handler(),
	})

	errChan := make(chan error)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func(errs chan error) {
		err := apiService.ListenAndServe()
		if err != nil {
			errs <- errors.Wrap(err, "API listen and serve error")
		}
	}(errChan)
	log.Printf("client-api service started on %s", apiAddr)

	select {
	case s := <-signalChan:
		log.Printf("captured exit signal %s", s)
		return nil
	case err = <-errChan:
		return err
	}

}
