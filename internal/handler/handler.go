package handler

import (
	"net/http"

	"github.com/PirateWindows/DriverMockup/internal/service"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	service.CreateUserFromRequest(c)
}
