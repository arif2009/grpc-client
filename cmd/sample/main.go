package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/synspective/grpc-client/cmd/sample/handlers"
	"github.com/synspective/grpc-client/pkg/core"
	gruidae "github.com/synspective/grpc-client/pkg/logger"
	"github.com/synspective/grpc-client/pkg/notification"
)

const serviceName = "grpc-client"
const version = "local"

func main() {
	// # Initialize app. This will fetch all sensitive
	// information from secret manager
	if err := core.Init(serviceName); err != nil {
		panic(err)
	}

	// # Connect to dababase
	// if err := models.Connect(); err != nil {
	// 	panic(err)
	// }

	// # Connect to ShapeDB
	// Make sure you have port-forward to the service if you are using this.
	// if err := shapedb.Connect(); err != nil {
	// 	panic(err)
	// }

	// Initialize logger
	gruidae.Init(serviceName, version)

	// Connect to Citadel.
	// Commented out since we're using a simple auth for demo purpose, but make sure
	// this is enabled when using with an actual service.
	// if err := citadel.Dial(); err != nil {
	// 	panic(err)
	// }

	// Connect to the citadel usermgmt
	if err := notification.Connect(); err != nil {
		panic(err)
	}

	app := fiber.New()

	// Customizable port number
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// You will most likely need to defined CORS settings as REST API
	// are usually used by web applications
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("http://localhost:%s", port),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Fiber does not add access log by default so let's add it.
	// This is a MUST HAVE.
	app.Use(logger.New())

	// Always register a handler on root that simply returns 200 OK.
	// This is for health-check purposes.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Always add version number as a prefix.
	// This will help in the case a major change needed on the API
	// without breaking the UI once the updated API is deployed.
	v1 := app.Group("v1")

	// Register each handlers
	handlers.SetupNotificationHandlers(v1)

	// Listen and serve
	gruidae.Client.Error(app.Listen(fmt.Sprintf(":%s", port)))
}
