package main

import (
	"github.com/emperorsixpacks/dailbot/src/internal/services"
	"github.com/emperorsixpacks/dailbot/src/internal/services/airtable"
	"github.com/emperorsixpacks/dailbot/src/pkg/logger"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/emperorsixpacks/dailbot/src/server"
)

func returnServices(config utils.AppSettings) map[string]services.Service {
	airtableService := airtable.NewAirtableSerice(config)
	return map[string]services.Service{
		"airtable": airtableService,
	}
}

func main() {
	appSettings := utils.GetConfig()
	logger.NewDefaultLogger()
  logger.DefaultLogger.Info("Starting server...")
	server := server.NewFiberServer(appSettings, returnServices(appSettings))
	server.Start()
}
