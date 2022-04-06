package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DogsDelete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	if err := dogsService.Delete(c.Request.Context(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": id.String(),
	})
}

func DogsCount(c *gin.Context) {

}

func DogsGet(c *gin.Context) {

}

func DogsGetOne(c *gin.Context) {

}

func DogsPost(c *gin.Context) {

}

func DogsPut(c *gin.Context) {

}
