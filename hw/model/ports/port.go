package ports

import (
	pb "github.com/karolhrdina/misc/hw/pb.go"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// Port is model for port from input file
type Port struct {
	Key         string
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

// FillProtobuf fills protobuf Port's fields with those fields
// of Port that are non-nil.
func (p *Port) FillProtobuf(protoPort *pb.Port) {
	protoPort.Id = p.Key

	if p.Name != nil {
		protoPort.Name = &wrapperspb.StringValue{Value: *p.Name}
	}

	if p.City != nil {
		protoPort.City = &wrapperspb.StringValue{Value: *p.City}
	}

	if p.Country != nil {
		protoPort.Country = &wrapperspb.StringValue{Value: *p.Country}
	}

	if p.Province != nil {
		protoPort.Province = &wrapperspb.StringValue{Value: *p.Province}
	}

	if p.Timezone != nil {
		protoPort.Timezone = &wrapperspb.StringValue{Value: *p.Timezone}
	}

	if p.Code != nil {
		protoPort.Code = &wrapperspb.StringValue{Value: *p.Code}
	}

	if p.Alias != nil {
		protoPort.Alias = &pb.RepeatedString{Value: *p.Alias}
	}

	if p.Regions != nil {
		protoPort.Regions = &pb.RepeatedString{Value: *p.Regions}
	}

	if p.Unlocs != nil {
		protoPort.Unlocs = &pb.RepeatedString{Value: *p.Unlocs}
	}
	if p.Coordinates != nil && len(*p.Coordinates) >= 2 {
		protoPort.Coordinates = &pb.Coordinates{
			X: (*p.Coordinates)[0],
			Y: (*p.Coordinates)[1],
		}
	}
}

// FromProtobuf returns new `Port` initialized with values
// from protobuf's `protoPort`.
func FromProtobuf(protoPort *pb.Port) Port {
	rv := Port{Key: protoPort.Id}

	if protoPort.Name != nil {
		rv.Name = &protoPort.Name.Value
	}
	if protoPort.City != nil {
		rv.City = &protoPort.City.Value
	}
	if protoPort.Country != nil {
		rv.Country = &protoPort.Country.Value
	}
	if protoPort.Province != nil {
		rv.Province = &protoPort.Province.Value
	}
	if protoPort.Timezone != nil {
		rv.Timezone = &protoPort.Timezone.Value
	}
	if protoPort.Code != nil {
		rv.Code = &protoPort.Code.Value
	}
	if protoPort.Alias != nil {
		rv.Alias = &protoPort.Alias.Value
	}
	if protoPort.Regions != nil {
		rv.Regions = &protoPort.Regions.Value
	}
	if protoPort.Coordinates != nil {
		x := make([]float32, 2)
		x[0] = protoPort.Coordinates.X
		x[1] = protoPort.Coordinates.Y
		rv.Coordinates = &x
	}
	if protoPort.Unlocs != nil {
		rv.Unlocs = &protoPort.Unlocs.Value
	}

	return rv
}
