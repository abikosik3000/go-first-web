package webhandler

import (
	"github.com/kataras/iris/v12"
)

type HealthCheker interface {
	Application() string
}

func (webHandler WebHanler) handlerHealthChekApplication(ctx iris.Context) {
	chekResult := webHandler.repositories.HealthCheker.Application()
	ctx.JSON(chekResult)
}
