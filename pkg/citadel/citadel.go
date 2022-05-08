package citadel

import (
	"context"
	"os"
	"time"

	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn

func Dial() error {
	serverAddr := os.Getenv("SYNS_CITADEL_STS_ADDRESS")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure()) // ! this needs to be a TLS configuration
	opts = append(opts, grpc.WithBlock())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, opts...)
	if err != nil {
		return err
	}

	Conn = conn

	return nil
}
