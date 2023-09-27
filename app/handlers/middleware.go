package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (h *Handler) middleware(c *gin.Context) {

	slog.Debug("middleware before", slog.String("method", c.Request.Method))

	// do before handler

	c.Next()

	slog.Debug("middleware after", slog.String("method", c.Request.Method))

	// do after handler
}
