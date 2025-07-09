package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/", showMain)
	r.GET("/esports", showEsports)
	r.GET("/lacrosse", showLacrosse)
	r.GET("/programming", showProgramming)
	//r.GET("/users", handler.GetUsers)
	//r.POST("/users", handler.CreateUser)
	return r
}

func showMain(c *gin.Context) {
	c.HTML(http.StatusOK, "main1.html", nil)
}

func showEsports(c *gin.Context) {
	c.HTML(http.StatusOK, "esports.html", nil)
}

func showLacrosse(c *gin.Context) {
	c.HTML(http.StatusOK, "lacrosse.html", nil)
}

func showProgramming(c *gin.Context) {
	c.HTML(http.StatusOK, "programming.html", nil)
}
