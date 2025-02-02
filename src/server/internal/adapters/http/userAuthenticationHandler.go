package handlers

import (
	"github.com/emperorsixpacks/dailbot/config"
	"github.com/gofiber/fiber/v2"
)

type userAuthenticationHandler struct {
	appConfig config.AppSettings
	app       *fiber.App
}

func (t *userAuthenticationHandler) Handle() fiber.Router {
	router := t.app.Group("/auth")
	router.Post("/airtable", t.airtableOAuthHandler)
	return router
}

func (t *userAuthenticationHandler) airtableOAuthHandler(c *fiber.Ctx) error {
	return nil
}

func UserAuthenticationHandler(appConfig config.AppSettings, app *fiber.App) *userAuthenticationHandler {
	return &userAuthenticationHandler{appConfig, app}
}
