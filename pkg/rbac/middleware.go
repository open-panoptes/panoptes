package rbac

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rancher/opni-monitoring/pkg/core"
)

type middleware struct {
	provider Provider
	codec    Codec
}

func (m *middleware) Handle(c *fiber.Ctx) error {
	userID := c.Locals(UserIDKey)
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	clusters, err := m.provider.SubjectAccess(context.Background(), &core.SubjectAccessRequest{
		Subject: userID.(string),
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if len(clusters.Items) == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	ids := make([]string, len(clusters.Items))
	for i, cluster := range clusters.Items {
		ids[i] = cluster.Id
	}
	c.Request().Header.Set(m.codec.Key(), m.codec.Encode(ids))
	return c.Next()
}

func NewMiddleware(provider Provider, codec Codec) func(*fiber.Ctx) error {
	mw := &middleware{
		provider: provider,
		codec:    codec,
	}
	return mw.Handle
}
