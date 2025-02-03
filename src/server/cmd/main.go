package main

import (
	"github.com/emperorsixpacks/dailbot/internal/services/airtable"
	"github.com/emperorsixpacks/dailbot/pkg/utils"
)

func main() {
	appSettings := utils.GetConfig()
	airtable.NewAirtableSerice(appSettings.Services.Airtable).Authorise()
	//server := server.NewFiberServer(appSettings)
	//	server.Start()
}
