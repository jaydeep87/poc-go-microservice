package controllers

import (
	"fmt"
	// json library

	"github.com/jaydeep87/poc-go-microservice/src/mgo-microservices/models"

	"github.com/gin-gonic/gin"
)

var userModel = new(models.UserModel)

type UserController struct{}

func (user *UserController) Create(c *gin.Context) {
	var data models.User

	if c.BindJSON(&data) != nil {
		c.JSON(402, gin.H{"statusCode": 402, "statusMessage": "Invalid form1", "form": data})
		c.Abort()
		return
	}

	err := userModel.Create(data)
	if err != nil {
		c.JSON(402, gin.H{"statusCode": 402, "statusMessage": "User could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"statusCode": 200, "statusMessage": "User Created"})
}

func (user *UserController) Find(c *gin.Context) {
	list, err := userModel.Find()
	if err != nil {
		c.JSON(402, gin.H{"statusCode": 402, "statusMessage": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (user *UserController) Get(c *gin.Context) {
	id := c.Param("id")
	profile, err := userModel.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"statusCode": 404, "statusMessage": "User not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": profile})
	}
}

func (user *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var data models.User

	if c.BindJSON(&data) != nil {
		c.JSON(402, gin.H{"statusCode": 402, "statusMessage": "Invalid Parameters"})
		c.Abort()
		return
	}
	err := userModel.Update(id, data)
	if err != nil {
		c.JSON(402, gin.H{"statusCode": 402, "statusMessage": "User Could Not Be Updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"statusCode": 200, "statusMessage": "User Updated"})
}

func (user *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := userModel.Delete(id)
	if err != nil {
		c.JSON(402, gin.H{"statusCode": 402, "statusMessage": "User Could Not Be Deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"statusCode": 200, "statusMessage": "User Deleted"})
}
