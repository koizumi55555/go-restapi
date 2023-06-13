package main

import (
	"koizumi55555/go-restapi/src"
	"koizumi55555/go-restapi/src/controller"
	"koizumi55555/go-restapi/src/infra"
	"koizumi55555/go-restapi/src/usecase/postgres"
	"log"
)

func main() {

	db, err := infra.Connect()
	if err != nil {
		log.Fatal(err)
	}

	p := postgres.New(db)
	uc := controller.NewUserController(p)

	engine := src.Server(uc)
	engine.Run(":8080")

}
