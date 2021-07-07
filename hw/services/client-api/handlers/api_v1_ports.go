package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/karolhrdina/misc/hw/model/ports"

	pb "github.com/karolhrdina/misc/hw/pb.go"
)

type PortsV1 struct {
}

func NewV1PortsImportAPI() *PortsV1 {
	return &PortsV1{}
}

func (s *PortsV1) HandleImport(storer ports.Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parser := ports.NewPortParser(r.Body)

		for parser.Next() {
			var port ports.Port
			if err := parser.Scan(&port); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			protoPort := &pb.Port{}
			port.FillProtobuf(protoPort)

			if err := storer.Store(r.Context(), protoPort); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func (s *PortsV1) HandleList(storer ports.Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		list, err := storer.List(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(w)

		for _, port := range list {
			err = encoder.Encode(ports.FromProtobuf(port))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
