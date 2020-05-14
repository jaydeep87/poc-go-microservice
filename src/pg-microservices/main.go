package main

import (
	"log"
	"os"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/op/go-logging"
	// "github.com/jaydeep87/poc-go-microservice/src/pg-microservices/logger"
	config "github.com/jaydeep87/poc-go-microservice/src/pg-microservices/config"
	routes "github.com/jaydeep87/poc-go-microservice/src/pg-microservices/routes"
)



func main() {
	path := "logs"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0666)
	}
	today := time.Now()
	ds := today.Format("2006-01-02")
	logpath:= "logs/" + ds + ".log"
	fmt.Println(logpath)
	   var _, err1 = os.Create(logpath)

   if err1 != nil {
      panic(err1)
   }
	f, err := os.OpenFile(logpath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	
	log.SetOutput(f)
	log.Println("This is a test log entry");

	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8082"))
}
