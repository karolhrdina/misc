package main

import (
	"encoding/json"
	"io"
	"net/http"

	pb "github.com/karolhrdina/misc/hw/pb.go"
	"google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// one port
type item struct {
	Key         *string    `json:"key"`
	Name        *string    `json:"name,omitempty"`
	City        *string    `json:"city,omitempty"`
	Country     *string    `json:"country,omitempty"`
	Alias       *[]string  `json:"alias,omitempty"`
	Regions     *[]string  `json:"regions,omitempty"`
	Coordinates *[]float32 `json:"coordinates,omitempty"`
	Province    *string    `json:"province,omitempty"`
	Timezone    *string    `json:"timezone,omitempty"`
	Unlocs      *[]string  `json:"unlocs,omitempty"`
	Code        *string    `json:"code,omitempty"`
}

type v1Ports struct {
	portDomainGRPC pb.PortdomainClient
}

func NewV1PortsImportAPI(portDomainClient pb.PortdomainClient) *v1Ports {
	return &v1Ports{
		portDomainGRPC: portDomainClient,
	}
}

func (s *v1Ports) handleImport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)

		// read open bracket
		t, err := dec.Token()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bracket, ok := t.(json.Delim)
		if !ok {
			http.Error(w, "json object exptected", http.StatusBadRequest)
			return
		}
		if bracket != '{' {
			http.Error(w, "json object expected", http.StatusBadRequest)
			return
		}

		// while the array contains values
		for dec.More() {

			t, err = dec.Token()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			id, ok := t.(string)
			if !ok {
				continue
			}
			port := &pb.Port{
				Id: id,
			}

			var m item
			// decode port
			err := dec.Decode(&m)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			setPBPortFields(port, m)
			_, err = s.portDomainGRPC.Snapshot(r.Context(), port)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// we can ignore reading closing bracket, since the data has already been sent
		// to port-domain for storage
	}
}

func (s *v1Ports) handleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stream, err := s.portDomainGRPC.List(r.Context(), &emptypb.Empty{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(w)

		for {
			port, err := stream.Recv()
			if err == io.EOF {
				break
			}
			// handle error
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = encoder.Encode(convertPBPortToItem(port))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

// convertPBPortToItem converts models.Port to protobuf's Port
func convertPBPortToItem(p *pb.Port) item {
	rv := item{Key: &p.Id}

	if p.Name != nil {
		rv.Name = &p.Name.Value
	}
	if p.City != nil {
		rv.City = &p.City.Value
	}
	if p.Country != nil {
		rv.Country = &p.Country.Value
	}
	if p.Province != nil {
		rv.Province = &p.Province.Value
	}
	if p.Timezone != nil {
		rv.Timezone = &p.Timezone.Value
	}
	if p.Code != nil {
		rv.Code = &p.Code.Value
	}
	if p.Alias != nil {
		rv.Alias = &p.Alias.Value
	}
	if p.Regions != nil {
		rv.Regions = &p.Regions.Value
	}
	if p.Coordinates != nil {
		x := make([]float32, 2)
		x[0] = p.Coordinates.X
		x[1] = p.Coordinates.Y
		rv.Coordinates = &x
	}
	if p.Unlocs != nil {
		rv.Unlocs = &p.Unlocs.Value
	}

	return rv
}

// setPBPortFields fills protobuf's Port fields
// with non-nil models.Port's fields.
func setPBPortFields(p *pb.Port, m item) {
	if m.Name != nil {
		p.Name = &wrapperspb.StringValue{Value: *m.Name}
	}

	if m.City != nil {
		p.City = &wrapperspb.StringValue{Value: *m.City}
	}

	if m.Country != nil {
		p.Country = &wrapperspb.StringValue{Value: *m.Country}
	}

	if m.Province != nil {
		p.Province = &wrapperspb.StringValue{Value: *m.Province}
	}

	if m.Timezone != nil {
		p.Timezone = &wrapperspb.StringValue{Value: *m.Timezone}
	}

	if m.Code != nil {
		p.Code = &wrapperspb.StringValue{Value: *m.Code}
	}

	if m.Alias != nil {
		p.Alias = &pb.RepeatedString{Value: *m.Alias}
	}

	if m.Regions != nil {
		p.Regions = &pb.RepeatedString{Value: *m.Regions}
	}

	if m.Unlocs != nil {
		p.Unlocs = &pb.RepeatedString{Value: *m.Unlocs}
	}

	if m.Coordinates != nil && len(*m.Coordinates) >= 2 {
		p.Coordinates = &pb.Coordinates{
			X: (*m.Coordinates)[0],
			Y: (*m.Coordinates)[1],
		}
	}
}
