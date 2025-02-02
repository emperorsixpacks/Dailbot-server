package goFiberHanders

import (
	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type webHookHandler struct {
	appConfig utils.AppSettings
}

func (w *webHookHandler) Handle() *fiber.App {

	router := fiber.New()
	return router
}

func WebHookHandler(appConfig utils.AppSettings) *webHookHandler {
	return &webHookHandler{appConfig}
}

func handleWebhook(c *fiber.Ctx) error {
	return nil
}
