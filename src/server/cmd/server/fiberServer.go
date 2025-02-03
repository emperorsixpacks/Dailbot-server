package server

import (
	"fmt"
	"path"
	"time"

	goFiberHanders "github.com/emperorsixpacks/dailbot/internal/adapters/http/FiberHandler"
	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// TODO pass the template file dir directly from the appsettings, also look into cdns for prod
func NewFiberServer(appConfig utils.AppSettings) *fiberServer {
	basePath, _ := utils.GetBasePath()
	templesDir := path.Join(basePath, "public/pages")
	staticDir := path.Join(basePath, "dist")
	templateEngine := html.New(templesDir, ".html")
	app := fiber.New(fiber.Config{
		AppName: appConfig.Server.Name,
		Views:   templateEngine,
	})
	app.Static("/static", staticDir, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})
	return &fiberServer{appConfig, app}
}

type fiberServer struct {
	appConfig utils.AppSettings
	app       *fiber.App
}

// TODO maybe add an initializer method

// TODO look into moving temples dir and static files path to appsettings
func (f *fiberServer) Start() {
	goFiberHanders.AirtableAuthnHandler(f.appConfig.Services.Airtable, f.app).Handle()
	if err := f.app.Listen(f.appConfig.Server.Port); err != nil {
		fmt.Println(err)
		// TODO log error
		return
	}
}
