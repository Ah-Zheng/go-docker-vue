package threekingdoms

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	server := gin.New()
	server.Use(static.Serve("/", static.LocalFile("./public/threekingdoms/", false)))

	// api := server.Group("/api")
	// {
	// 	api.GET("/generate", guestNumber.GenerateNumber)
	// 	api.POST("/guest", guestNumber.Guest)
	// }

	return server
}
