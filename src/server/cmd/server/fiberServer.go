package server

import (
	"fmt"
	"path"

	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// TODO pass the template file dir directly from the appsettings, also look into cdns for prod
func NewFiberServer(appConfig utils.AppSettings) *fiberServer {
	basePath, _ := utils.GetBasePath()
  fmt.Println(basePath)
	templesDir := path.Join(basePath, "public/pages")
	templateEngine := html.New(templesDir, ".html")
	app := fiber.New(fiber.Config{
		Views: templateEngine,
	})
	return &fiberServer{appConfig, app}
}

type fiberServer struct {
	appConfig utils.AppSettings
	app       *fiber.App
}

// TODO look into moving temples dir and static files path to appsettings
func (f *fiberServer) Start() {

	if err := f.app.Listen(f.appConfig.Server.Port); err != nil {
		fmt.Println(err)
		// TODO log error
		return
	}
}
