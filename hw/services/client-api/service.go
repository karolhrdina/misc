package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Service is a container to hold all dependencies of client-api app.
type Service struct {
	portsImport *v1Ports
}

// NewService instantiates Service struct.
func NewService(
	portsImport *v1Ports) *Service {
	return &Service{
		portsImport: portsImport,
	}
}

// Handler returns *mux.Router which conforms to the http.Handler interface.
func (c *Service) Handler() http.Handler {
	// Note:
	//  In real-life app, a bit more would be going-on in here
	//  * readiness check for dependent gRPC port-domain service
	//  * /ready /live endpoints
	//  * middleware for error logging, metrics, panic recovery, maybe auth and so on...
	//  * tracing ...

	apiRouter := mux.NewRouter().StrictSlash(true)
	v1Router := apiRouter.PathPrefix("/v1").Subrouter()
	v1Router.Methods(http.MethodPost).Path("/ports/import").HandlerFunc(c.portsImport.handleImport())
	v1Router.Methods(http.MethodGet).Path("/ports").HandlerFunc(c.portsImport.handleList())

	return apiRouter
}

// Close closes all dependencies
func (s *Service) Close() error {
	return nil
}
