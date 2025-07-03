package service

import (
	"github.com/PirateWindows/DriverMockup/internal/model"
	"github.com/PirateWindows/DriverMockup/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetAllUsers() []model.User {
	return repository.GetUsers()
}

func CreateUserFromRequest(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	repository.AddUser(user)
	c.JSON(201, user)
}
