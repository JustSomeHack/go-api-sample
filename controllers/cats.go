package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CatsDelete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	if err := catsService.Delete(c.Request.Context(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": id.String(),
	})
}

func CatsCount(c *gin.Context) {

}

func CatsGet(c *gin.Context) {

}

func CatsGetOne(c *gin.Context) {

}

func CatsPost(c *gin.Context) {

}

func CatsPut(c *gin.Context) {

}
