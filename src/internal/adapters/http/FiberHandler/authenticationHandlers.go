package goFiberHanders

import (
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type oauthService interface {
	ConnectURL() string
}

type authHandlers struct {
	appConfig utils.AirtableSettings
	app       *fiber.App
	services  map[string]*oauthService
}

func (a authHandlers) getService(name string) *oauthService {
	for k, v := range a.services {
		if k == name {
			return v
		}
	}
	return nil
}

func (a *authHandlers) Handle() *fiber.Router {
	router := a.app.Group("/auth")
	router.Get("/auth/:service_name", a.authenticeSerivice)
	return &router
}

func (a *authHandlers) authenticeSerivice(ctx *fiber.Ctx) error {
	return nil
}

func (a *authHandlers) callback(ctx *fiber.Ctx) error {
	return nil
}

func NewauthHandler(appConfig utils.AirtableSettings, app *fiber.App, service map[string]*oauthService) *authHandlers {
	return &authHandlers{appConfig, app, service}

}
