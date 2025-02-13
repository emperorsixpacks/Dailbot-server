package main

import (
	"github.com/emperorsixpacks/dailbot/src/pkg/logger"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/emperorsixpacks/dailbot/src/server"
)

func main() {
	appSettings := utils.GetConfig()
	logger.NewDefaultLogger()
	server := server.NewFiberServer(appSettings)
	server.Start()
}
