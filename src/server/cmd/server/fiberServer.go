package server

import (
	"fmt"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewFiberServer(appConfig config.AppSettings) *fiberServer {
	basePath, _ := utils.GetBasePath()
	templesDir := path.Join(basePath, "public/pages")
	templateEngine := html.New(templesDir, ".html")
	app := fiber.New(fiber.Config{
		Views: templateEngine,
	})
	return &fiberServer{appConfig, app}
}

type fiberServer struct {
	appConfig config.AppSettings
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
