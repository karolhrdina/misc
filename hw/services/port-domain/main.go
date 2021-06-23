package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/karolhrdina/misc/hw/pkg/env"

	"github.com/pkg/errors"
)

var (
	grpcAddr = env.GetVar("GRPC_ADDR", ":9082")
)

func main() {
	if err := runGRPC(); err != nil {
		log.Fatalf("run error: %s", err.Error())
		os.Exit(1)
	}
}

func runGRPC() error {
	service, err := initializeService()
	if err != nil {
		return errors.Wrap(err, "unable to initialize service")
	}
	defer func() {
		errC := service.Close()
		if errC != nil {
			log.Printf("error closing service: %s", errC.Error())
		}
	}()

	errChan := make(chan error)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	server, err := service.GRPC()
	if err != nil {
		return errors.Wrap(err, "error initializing grpc")
	}
	go func(errs chan error) {
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			errs <- errors.Wrap(err, "grpc `net.Listen()` failed")
			return
		}
		err = server.Serve(lis)
		if err != nil {
			errs <- errors.Wrap(err, "grpc `Server()` error")
		}
	}(errChan)
	log.Printf("port-domain gRPC service started on %s", grpcAddr)

	select {
	case s := <-signalChan:
		log.Printf("captured exit signal %s", s)
		server.GracefulStop()
		return nil
	case err = <-errChan:
		return err
	}
}
