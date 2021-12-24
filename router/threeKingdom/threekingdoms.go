package threekingdoms

import (
	"go-docker-vue/db"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	server := gin.New()
	server.Use(static.Serve("/", static.LocalFile("./public/threekingdoms/", false)))

	db.SqlConn()

	// api := server.Group("/api")
	// {
	// 	api.GET("/tactics", crawler.GetTacticInfo())
	// }

	return server
}
