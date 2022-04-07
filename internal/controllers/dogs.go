package controllers

import (
	"fmt"
	"net/http"

	"github.com/JustSomeHack/go-api-sample/internal/models"
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
	dogs, err := dogsService.Get(c.Request.Context(), nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}
	c.JSON(http.StatusOK, dogs)
}

func DogsGetOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	dog, err := dogsService.GetOne(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}
	c.JSON(http.StatusOK, dog)
}

func DogsPost(c *gin.Context) {
	dog := new(models.Dog)
	if c.ShouldBind(dog) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
		})
		return
	}

	if dog.ID == uuid.Nil {
		dog.ID = uuid.New()
	}

	id, err := dogsService.Add(c.Request.Context(), dog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("id %s created", id.String()),
	})
}

func DogsPut(c *gin.Context) {
	dog := new(models.Dog)
	if c.ShouldBind(&dog) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	if err := dogsService.Update(c.Request.Context(), id, dog); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": fmt.Sprintf("id %s updated", id.String()),
	})
}
