package goFiberHanders

import (
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type servicesHandler struct {
	appConfig utils.AirtableSettings
	app       *fiber.App
}

func (t *servicesHandler) Handle() fiber.Router {
	router := t.app.Group("/services")
	router.Get("/", t.getIntegrations)
	//	router.Post("/airtable", t.airtableOAuthHandler)
	return router
}

func (t *servicesHandler) getIntegrations(ctx *fiber.Ctx) error {
	ctx.Accepts("html")
	return ctx.Render("addService", nil)
}

func AirtableAuthnHandler(appConfig utils.AirtableSettings, app *fiber.App) *servicesHandler {
	return &servicesHandler{appConfig, app}
}
