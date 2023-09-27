package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) getRecord(c *gin.Context) {
	slog.Info("GET", slog.String("id", c.Param("id")))

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) postRecord(c *gin.Context) {
	slog.Info("POST")

	c.JSON(http.StatusOK, nil)
}
