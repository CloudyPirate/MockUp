package router

import (
	"net/http"

	"github.com/PirateWindows/DriverMockup/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/", showMain)
	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)
	return r
}

func showMain(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", nil)
}
