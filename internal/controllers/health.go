package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Gets the status of the server
// @Description gets the status of the server
// @Produce  json
// @Success      204   {string}  string  "answer"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /health [get]
func HealthGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}