package goFiberHanders

import (
	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type airtableAuthHandler struct {
	appConfig utils.AirtableSettings
	app       *fiber.App
}

func (t *airtableAuthHandler) Handle() fiber.Router {
	router := t.app.Group("/auth")
	//	router.Get("/airtable", t.airtableOAuthHandler)
	//	router.Post("/airtable/authenticate", t.airtableOAuthHandler)
	return router
}

func AirtableAuthnHandler(appConfig utils.AirtableSettings, app *fiber.App) *airtableAuthHandler {
	return &airtableAuthHandler{appConfig, app}
}
