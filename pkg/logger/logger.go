package logger

import (
	"context"
	"log"
	"os"

	"github.com/synspective/gruidae-logging-go/logging"
)

var Client *logging.Logger

func Init(serviceName string, version string) {
	ctx := context.Background()

	stackdriverHandler, err := logging.InitStackdriverHandler(ctx, os.Getenv("PROJECT_ID"), serviceName, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		panic(err)
	}

	Client = logging.InitLogger(serviceName, version, &logging.StreamHandler{}, stackdriverHandler)
}
