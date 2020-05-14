package main

import (
	"net/http"
	"log"
    "os"

	"github.com/jaydeep87/poc-go-microservice/src/mgo-microservices/controllers"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

var (
	// outfile, _ = os.Create("logs/my.log") // update path for your needs
	outfile, err = os.OpenFile("logs/my.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    l      = log.New(outfile, "", 0)
)

const logPath = "logs/mylog.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {

	flag.Parse()

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
	  logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()
	
	// Log to system log and a log file, Info logs don't write to stdout.
	loggerOne := logger.Init("LoggerExample", false, true, lf)
	defer loggerOne.Close()
	// Don't to system log or a log file, Info logs write to stdout..
	loggerTwo := logger.Init("LoggerExample", true, false, ioutil.Discard)
	defer loggerTwo.Close()
	
	loggerOne.Info("This will log to the log file and the system log")
	loggerTwo.Info("This will only log to stdout")
	logger.Info("This is the same as using loggerOne")
	
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
