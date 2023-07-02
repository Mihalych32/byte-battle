package handler

import "github.com/gin-gonic/gin"

func StartNewServer(router *gin.Engine, port string) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}

	router.Run(port)
}
