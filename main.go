package main

import (
	"koizumi55555/go-restapi/src"
	"koizumi55555/go-restapi/src/controller"
	"koizumi55555/go-restapi/src/infra"
	"koizumi55555/go-restapi/src/usecase/postgres"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("src/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := infra.Connect()
	if err != nil {
		log.Fatal(err)
	}

	p := postgres.New(db)
	uc := controller.NewUserController(p)
	s := controller.NewServerController(
		strings.TrimSpace(os.Getenv("CLIENT_ID")),
		strings.TrimSpace(os.Getenv("CLIENT_SECRET")),
	)

	engine := src.Server(uc, s)
	engine.Run(":8080")

}
