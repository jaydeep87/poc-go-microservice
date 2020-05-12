package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "github.com/jaydeep87/poc-go-microservice/src/pg-microservices/config"
	routes "github.com/jaydeep87/poc-go-microservice/src/pg-microservices/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8082"))
}
