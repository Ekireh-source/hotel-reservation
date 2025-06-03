package main

import (
	"flag"

	"github.com/Ekireh-source/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)


func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the api server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/use/:id", api.HandleGetUser)
	app.Listen(*listenAddr)

}


