package main

import (
	"fist-web-go/internal/db/postgres"
	"fist-web-go/internal/repositories/healthchek"
	"fist-web-go/internal/webhandler"
	"fmt"
	"log"
)

func main() {
	fmt.Println("starting")

	db, err := makeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	repositories, err := makeRepositories(db)
	if err != nil {
		log.Fatal(err)
	}

	webHandler, _ := webhandler.New(8080, repositories)
	webHandler.Run()
}

func makeDatabase() (*postgres.Database, error) {
	return postgres.New(postgres.Config{
		Database: "db1",
		User:     "user",
		Password: "hackme",
		Host:     "localhost",
		Port:     5432,
	})
}

func makeRepositories(db *postgres.Database) (webhandler.Repositories, error) {
	return webhandler.Repositories{
		HealthCheker: healthchek.NewRepository(db),
	}, nil
}
