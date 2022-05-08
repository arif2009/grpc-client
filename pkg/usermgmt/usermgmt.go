package usermgmt

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	u "github.com/synspective/syns-citadel-v2/pkg/api/proto/services/usermgmt"
	"google.golang.org/grpc"
)

var Usermgmt u.UserManagementServiceClient

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

	Usermgmt = u.NewUserManagementServiceClient(conn)

	return nil
}
