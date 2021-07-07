package storer

import (
	"context"
	"database/sql"
	"strconv"

	pb "github.com/karolhrdina/misc/hw/pb.go"
	"github.com/karolhrdina/misc/hw/services/port-domain/sql/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/pkg/errors"
)

// Storage implements the ports.Storer interface
type PostgresStorer struct {
	db boil.ContextExecutor
}

// NewPostgresStorer returns initialized *Storage
func NewPostgresStorer(db *sql.DB) *PostgresStorer {
	return &PostgresStorer{
		db: db,
	}
}

func (s *PostgresStorer) Store(ctx context.Context, in *pb.Port) error {
	if in == nil {
		return errors.New("function parameter `port` is nil")
	}

	port, blacklist := convertMessageToModel(in)
	err := port.Upsert(ctx, s.db, true, []string{models.PortColumns.Key}, blacklist, boil.Infer())
	if err != nil {
		return errors.Wrapf(err, "storing port to database failed")
	}

	return nil
}

func (s *PostgresStorer) List(ctx context.Context) ([]*pb.Port, error) {
	ports, err := models.Ports().All(ctx, s.db)
	if err != nil {
		return nil, errors.Wrapf(err, "loading ports from database failed")
	}
	rv := make([]*pb.Port, 0, len(ports))
	for _, cpy := range ports {
		rv = append(rv, convertModelToMessage(*cpy))
	}
	return rv, nil
}

// convertMessageToModel returns new models.Port created from protobuf's Port
// and furthermore returns a list of columns that must not be updated, i.e.
// those which a corresponding protobuf's Port fields is nil
func convertMessageToModel(in *pb.Port) (models.Port, boil.Columns) {
	port := models.Port{Key: in.Id}
	blacklist := boil.Blacklist()

	if in.Name != nil {
		port.Name = null.StringFrom(in.Name.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Name)
	}

	if in.City != nil {
		port.City = null.StringFrom(in.City.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.City)
	}

	if in.Country != nil {
		port.Country = null.StringFrom(in.Country.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Country)
	}

	if in.Province != nil {
		port.Province = null.StringFrom(in.Province.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Province)
	}

	if in.Timezone != nil {
		port.Timezone = null.StringFrom(in.Timezone.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Timezone)
	}

	if in.Code != nil {
		port.Code = null.StringFrom(in.Code.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Code)
	}
	if in.Alias != nil {
		port.Alias = types.StringArray(in.Alias.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Alias)
	}

	if in.Regions != nil {
		port.Regions = types.StringArray(in.Regions.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Regions)
	}

	if in.Coordinates != nil {
		port.Xcoord = null.StringFrom(strconv.FormatFloat(float64(in.Coordinates.X), 'f', -1, 32))
		port.Ycoord = null.StringFrom(strconv.FormatFloat(float64(in.Coordinates.Y), 'f', -1, 32))
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Xcoord)
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Ycoord)
	}

	if in.Unlocs != nil {
		port.Unlocs = types.StringArray(in.Unlocs.Value)
	} else {
		blacklist.Cols = append(blacklist.Cols, models.PortColumns.Unlocs)
	}

	return port, blacklist
}

// convertModelToMessage returns new protobuf's Port created from models.Port
func convertModelToMessage(in models.Port) *pb.Port {
	port := &pb.Port{
		Id: in.Key,
	}

	if !in.Name.IsZero() {
		port.Name = &wrapperspb.StringValue{Value: in.Name.String}
	}

	if !in.City.IsZero() {
		port.City = &wrapperspb.StringValue{Value: in.City.String}
	}

	if !in.Country.IsZero() {
		port.Country = &wrapperspb.StringValue{Value: in.Country.String}
	}

	if !in.Province.IsZero() {
		port.Province = &wrapperspb.StringValue{Value: in.Province.String}
	}

	if !in.Timezone.IsZero() {
		port.Timezone = &wrapperspb.StringValue{Value: in.Timezone.String}
	}

	if !in.Code.IsZero() {
		port.Code = &wrapperspb.StringValue{Value: in.Code.String}
	}

	port.Alias = &pb.RepeatedString{Value: in.Alias}
	port.Regions = &pb.RepeatedString{Value: in.Regions}
	port.Unlocs = &pb.RepeatedString{Value: in.Unlocs}

	if !in.Xcoord.IsZero() && !in.Ycoord.IsZero() {
		xf, err := strconv.ParseFloat(in.Xcoord.String, 32)
		yf, err2 := strconv.ParseFloat(in.Ycoord.String, 32)
		if err == nil && err2 == nil {
			port.Coordinates = &pb.Coordinates{
				X: float32(xf),
				Y: float32(yf),
			}
		}
	}

	return port
}
