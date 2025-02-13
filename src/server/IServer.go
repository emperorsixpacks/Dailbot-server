package server

import handler "github.com/emperorsixpacks/dailbot/src/internal/adapters/http"

type Server interface {
	Start() error
	AddHandler(handler.Hander)
}
