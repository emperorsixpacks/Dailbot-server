package server

import handler "github.com/emperorsixpacks/dailbot/internal/adapters/http"

type Server interface {
	Start() error
	AddHandler(handler.Hander)
}
