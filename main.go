package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.StaticFS("/assets", http.Dir("./assets"))
	initializeRoutes()

	router.Run(":8081")
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])

	case "application/xml":
		c.XML(http.StatusOK, data["payload"])

	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
