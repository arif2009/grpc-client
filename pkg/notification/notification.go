package notification

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	nt "github.com/synspective/syns-citadel-v2/pkg/api/proto/services/notification"
	"google.golang.org/grpc"
)

var Notification nt.NotificationServiceClient

const dialTimeout = time.Second * 15

func Connect() error {
	host := os.Getenv("NOTIFY_HOST")
	port := os.Getenv("NOTIFY_PORT")
	if host == "" || port == "" {
		return errors.New("missing Notification env variables")
	}

	address := fmt.Sprintf("%s:%s", host, port)

	// Set up a connection to the server.
	ctx, ccl := context.WithTimeout(context.Background(), dialTimeout)
	defer ccl()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	Notification = nt.NewNotificationServiceClient(conn)

	return nil
}
