package goFiberHanders

import (
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
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
func WebHookHandler(appConfig utils.AirtableSettings, app *fiber.App) *webHookHandler {
	return &webHookHandler{appConfig, app}
}
