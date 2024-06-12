package routes

import (
	"github.com/gin-gonic/gin"
	"myapi/controllers"
)

func PingRoute(router *gin.RouterGroup) {
	auth := router.Group("/ping")
	{
		auth.GET(
			"",
			controllers.Ping,
		)
	}
}
