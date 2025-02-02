package server

import "github.com/emperorsixpacks/dailbot/config"

type Server interface {
	Start(config.AppSettings)
}
