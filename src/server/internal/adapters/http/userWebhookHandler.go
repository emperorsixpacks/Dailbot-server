package handlers

import (
	"github.com/emperorsixpacks/dailbot/config"
	"github.com/gofiber/fiber/v2"
)

type webHookHandler struct {
	appConfig config.AppSettings
}

func (w *webHookHandler) Handle() *fiber.App {

	router := fiber.New()
	return router
}

func WebHookHandler(appConfig config.AppSettings) *webHookHandler {
	return &webHookHandler{appConfig}
}

func handleWebhook(c *fiber.Ctx) error {
	return nil
}
