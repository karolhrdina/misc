package service

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/karolhrdina/misc/hw/model/ports"
	"github.com/karolhrdina/misc/hw/services/client-api/handlers"
)

// Service is a container to hold all dependencies of client-api app.
type Service struct {
	portsHandlers *handlers.PortsV1
	storer        ports.Storer
}

// InitializeServices provides new Service injected with dependencies
func InitializeService() (*Service, error) {
	return initializeService()
}

// NewService instantiates Service struct.
func NewService(
	portsHandlers *handlers.PortsV1,
	storer ports.Storer) *Service {
	return &Service{
		portsHandlers: portsHandlers,
		storer:        storer,
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
	v1Router.Methods(http.MethodPost).Path("/ports/import").HandlerFunc(c.portsHandlers.HandleImport(c.storer))
	v1Router.Methods(http.MethodGet).Path("/ports").HandlerFunc(c.portsHandlers.HandleList(c.storer))

	return apiRouter
}

// Close closes all dependencies
func (s *Service) Close() error {
	return nil
}
