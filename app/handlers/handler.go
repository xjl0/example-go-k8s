package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const basePathAPI = "/api/v1"

// Handler struct
type Handler struct {
}

// NewHandler constructor
func NewHandler() *Handler {
	return &Handler{}
}

// InitRoutes init http router
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))

	router.OPTIONS(basePathAPI+"/*any", func(c *gin.Context) {
		c.Status(204)
	})

	v1 := router.Group(basePathAPI, h.middleware)
	{

		v1.GET("/:id", h.getRecord)
		v1.POST("/", h.postRecord)

	}
	return router
}
