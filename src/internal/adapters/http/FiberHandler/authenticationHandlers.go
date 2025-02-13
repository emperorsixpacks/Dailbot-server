package goFiberHanders

import (
	"fmt"
	"net/url"

	"github.com/emperorsixpacks/dailbot/src/internal/services"
	"github.com/emperorsixpacks/dailbot/src/pkg/logger"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type authHandlers struct {
	appConfig utils.AirtableSettings
	app       *fiber.App
	services  map[string]services.Service
}

func (a authHandlers) getService(name string) services.Service {
	for k, v := range a.services {
		if k == name {
			return v
		}
	}
	return nil
}

func (a *authHandlers) Handle() *fiber.Router {
	router := a.app.Group("/auth")
	router.Get("/service=:service_name", a.authenticeSerivice)
	router.Get("/callback", a.callback)
	return &router
}

func (a *authHandlers) authenticeSerivice(ctx *fiber.Ctx) error {
	serviceName := ctx.Params("service_name")
	service := a.getService(serviceName)
	if service == nil {
		// return to error previous page with message
		logger.DefaultLogger.Error("Service not found: %v", serviceName)
		return nil
	}

	url, err := url.QueryUnescape(service.AuthURL())
	if err != nil {
		return nil
	}
	return ctx.Redirect(url)
}

func (a *authHandlers) callback(ctx *fiber.Ctx) error {
  fmt.Println(ctx.Queries())
	return nil
}

func NewAuthHandler(appConfig utils.AirtableSettings, app *fiber.App, service map[string]services.Service) *authHandlers {
	return &authHandlers{appConfig, app, service}

}
