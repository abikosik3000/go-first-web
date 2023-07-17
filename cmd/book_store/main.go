package main

import (
	"fist-web-go/internal/db/postgres"
	"fist-web-go/internal/repositories/healthchek"
	webhandler "fist-web-go/internal/web-handler"
	"fmt"
)

func main() {
	fmt.Println("starting")
	repositories := makeRepositories()
	webHandler, _ := webhandler.New(8080, repositories)
	webHandler.Run()
}

func makeDatabase() {
	postgres.New(postgres.Config{
		Database: "db1",
		User:     "user",
		Password: "hackme",
		Host:     "localhost",
		Port:     5432,
	})
}

func makeRepositories() webhandler.Repositories {
	return webhandler.Repositories{
		HealthCheker: healthchek.NewRepository(),
	}
}
