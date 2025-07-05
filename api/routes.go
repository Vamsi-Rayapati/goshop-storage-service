package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartbot/storage/api/presign"
	"github.com/smartbot/storage/middleware"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	uploadGroup := router.Group("/storage/api/v1")

	// authenticated routes
	uploadGroup.Use(middleware.Authenticate())
	{
		presign.RegisterRoutes(uploadGroup)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
		c.Abort()
	})
	return router

}
