package main

import (
	"github.com/emperorsixpacks/dailbot/cmd/server"
	"github.com/emperorsixpacks/dailbot/pkg/utils"
)

func main() {
	appSettings := utils.GetConfig()
	server := server.NewFiberServer(appSettings)
	server.Start()
}
