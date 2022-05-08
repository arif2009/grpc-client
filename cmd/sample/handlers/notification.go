package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/synspective/grpc-client/pkg/notification"
	"github.com/synspective/grpc-client/pkg/responder"
	upb "github.com/synspective/syns-citadel-v2/pkg/api/proto/services/notification"
)

func SetupNotificationHandlers(app fiber.Router) {
	// Separate route registration and the actual handler func
	//app.Get("orgntf/list", listNotification)
	//app.Get("orgntf/:id", oneNotification)
	app.Post("orgntf", createOrganization)
}

func createOrganization(c *fiber.Ctx) error {
	ctx, ccl := context.WithTimeout(context.Background(), grpcTimeout)
	defer ccl()
	req := &upb.CreateOrganizationNotificationRequest{
		ServerIdentityToken: os.Getenv("SYNS_CITADEL_SERVER_IDENTITY_TOKEN"),
		OrganizationId:      1,
		ServiceId:           1,
		Message:             "Test create organization notification-2",
		ResourceId:          1,
		ResourceType:        "test",
		ResourceSubtype:     "test-test",
	}
	res, err := notification.Notification.CreateOrganizationNotification(ctx, req)
	fmt.Println("Err:", err)
	if err != nil {
		return err
	}

	return responder.Send(c, fiber.Map{"data": res})
}

/*func listNotification(c *fiber.Ctx) error {
	ctx, ccl := context.WithTimeout(context.Background(), grpcTimeout)
	defer ccl()
	req := &upb.ListOrgNotificationsRequest{
		OrgId: 1,
	}
	res, err := notification.Notification.ListOrgNotifications(ctx, req)
	fmt.Println("Err:", err)
	if err != nil {
		return err
	}

	return responder.Send(c, fiber.Map{"data": res})
}

func oneNotification(c *fiber.Ctx) (err error) {
	var id int64
	if idstr := c.Params("id"); idstr != "" {
		id, err = strconv.ParseInt(idstr, 10, 32)
		if err != nil {
			return err
		}
	}
	ctx, ccl := context.WithTimeout(context.Background(), grpcTimeout)
	defer ccl()
	req := &upb.ListOrgNotificationsRequest{
		OrgId: int32(id),
	}
	res, err := notification.Notification.ListOrgNotifications(ctx, req)
	fmt.Println("Err:", err)
	if err != nil {
		return err
	}

	return responder.Send(c, fiber.Map{"data": res})
}*/
