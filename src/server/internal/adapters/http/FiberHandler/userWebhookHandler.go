package goFiberHanders

import (
	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type webHookHandler struct {
	appConfig utils.AirtableSettings
	app       *fiber.App
}

func (w *webHookHandler) Handle() *fiber.Router {

	router := w.app.Group("/webhook")
	router.Post(":user_webhhok_id")
	return &router
}

func (w *webHookHandler) webhook(ctx *fiber.Ctx) error {
	webhook_id := ctx.Params("user_webhook_id")
	// TODO this is our webhook that will listen for events
}

func WebHookHandler(appConfig utils.AirtableSettings, app *fiber.App) *webHookHandler {
	return &webHookHandler{appConfig, app}
}
