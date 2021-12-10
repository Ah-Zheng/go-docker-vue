package test

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	guestNumber "go-docker-vue/handler/guestNumber"
)

func SetUpRouter() *gin.Engine {
	server := gin.New()
	server.Use(static.Serve("/", static.LocalFile("./public/test/", false)))

	api := server.Group("/api")
	{
		api.GET("/generate", guestNumber.GenerateNumber)
		api.POST("/guest", guestNumber.Guest)
	}

	return server
}
