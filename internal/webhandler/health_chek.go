package webhandler

import (
	"log"

	"github.com/kataras/iris/v12"
)

type HealthCheker interface {
	Application() string
	Database() error
}

func (webHandler WebHanler) handlerHealthChekApplication(ctx iris.Context) {
	chekResult := webHandler.repositories.HealthCheker.Application()
	ctx.JSON(chekResult)
}

func (webHandler WebHanler) handlerHealthChekDatabase(ctx iris.Context) {
	err := webHandler.repositories.HealthCheker.Database()
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON("dbWorked")
}
