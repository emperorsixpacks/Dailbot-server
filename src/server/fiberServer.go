package server

import (
	"fmt"
	"path"

	goFiberHanders "github.com/emperorsixpacks/dailbot/src/internal/adapters/http/FiberHandler"
	"github.com/emperorsixpacks/dailbot/src/internal/services"
	"github.com/emperorsixpacks/dailbot/src/pkg/logger"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// TODO pass the template file dir directly from the appsettings, also look into cdns for prod
func NewFiberServer(appConfig utils.AppSettings, services map[string]services.Service) *fiberServer {
	basePath, _ := utils.GetBasePath()
	templesDir := path.Join(basePath, "public/pages")
	staticDir := path.Join(basePath, "dist")
	templateEngine := html.New(templesDir, ".html")

	templateEngine.AddFuncMap(map[string]interface{}{
		"PMGreeting":    getTimeOfDay,
		"UserFirstName": getFirstName,
	})
	app := fiber.New(fiber.Config{
		AppName: appConfig.Server.Name,
		Views:   templateEngine,
	})
	app.Static("/static", staticDir, fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		// CacheDuration: 10 * time.Second,
		//	MaxAge:        3600,
	})
	return &fiberServer{appConfig, app, services}
}

type fiberServer struct {
	appConfig utils.AppSettings
	app       *fiber.App
	services  map[string]services.Service
}

// TODO maybe add an initializer method

// TODO look into moving temples dir and static files path to appsettings
func (f *fiberServer) Start() {
	logger.DefaultLogger.Info("Adding handlers")
	goFiberHanders.NewAuthHandler(f.appConfig.Services.Airtable, f.app, f.services).Handle()
	goFiberHanders.NewIntegrationsHandler(f.appConfig.Services.Airtable, f.app).Handle()
	logger.DefaultLogger.Info("Done adding handlers")
	if err := f.app.Listen(fmt.Sprintf(":%d", f.appConfig.Server.Port)); err != nil {
		fmt.Println(err)
		// TODO log error
		return
	}
}
