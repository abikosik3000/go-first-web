package webhandler

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

type WebHanler struct {
	port         int
	app          *iris.Application
	repositories Repositories
}

type Repositories struct {
	HealthCheker HealthCheker
}

func New(port int, repositories Repositories) (*WebHanler, error) {
	app := iris.New()
	handler := WebHanler{
		port:         port,
		app:          app,
		repositories: repositories,
	}
	handler.setRoute()
	return &handler, nil
}

func (webHandler WebHanler) setRoute() error {
	webHandler.app.Get("/health_chek/application", webHandler.handlerHealthChekApplication)
	webHandler.app.Get("/health_chek/database", webHandler.handlerHealthChekDatabase)
	return nil
}

func (webHandler WebHanler) Run() error {
	port := fmt.Sprintf(":%d", webHandler.port)
	webHandler.app.Listen(port)
	return nil
}
