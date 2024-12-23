package main

import (
	"log"

	"github.com/gofiber/fiber"
)


func main(){

	app := fiber.New()

	err:=app.Listen(":5000")
	if err !=nil {
		log.Fatalf("Error",err)
	}
}