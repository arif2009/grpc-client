package shapedb

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/paulmach/orb/geojson"
	shpgrpc "github.com/synspective/country-city-area/models/pb"
	"google.golang.org/grpc"
)

var Client shpgrpc.ShapedbClient

const dialTimeout = time.Second * 15

// ShapeDB returns the GeoJSON as bytes but we want to send back
// as actual GeoJSON so let's override it with our own struct.
type Area struct {
	*shpgrpc.AOI
	GeoJSON *geojson.FeatureCollection `json:"geojson,omitempty"`
}

func Connect() error {
	host := os.Getenv("SHAPEDB_HOST")
	port := os.Getenv("SHAPEDB_PORT")
	if host == "" || port == "" {
		return errors.New("missing ShapeDB env variables")
	}

	address := fmt.Sprintf("%s:%s", host, port)

	// Set up a connection to the server.
	ctx, ccl := context.WithTimeout(context.Background(), dialTimeout)
	defer ccl()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	Client = shpgrpc.NewShapedbClient(conn)

	return nil
}
