package main

import (
	"net/http"

	"github.com/poc-go-microservice/src/mgo-microservices/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	// or
	// router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		user := new(controllers.UserController)
		v1.POST("/users", user.Create)
		v1.GET("/users", user.Find)
		v1.GET("/users/:id", user.Get)
		v1.PUT("/users/:id", user.Update)
		v1.DELETE("/users/:id", user.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":8081")
}
