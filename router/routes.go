package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/healthcheck", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Metrics",
			})
		})

		v1.POST("/metrics", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Metrics",
			})
		})

	}

}
