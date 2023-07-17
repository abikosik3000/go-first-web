package main

import (
	"fist-web-go/internal/repositories/healthchek"
	webhandler "fist-web-go/internal/web-handler"
	"fmt"
)

func main() {
	fmt.Println("starting")
	repositories := makeDependens()
	webHandler, _ := webhandler.New(8080, repositories)
	webHandler.Run()
}

func makeDependens() webhandler.Repositories {
	return webhandler.Repositories{
		HealthCheker: healthchek.NewRepository(),
	}
}
