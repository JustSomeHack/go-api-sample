package controllers

import (
	"fmt"
	"net/http"

	"github.com/JustSomeHack/go-api-sample/internal/models"

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
	cats, err := catsService.Get(c.Request.Context(), nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}
	c.JSON(http.StatusOK, cats)
}

func CatsGetOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	cat, err := catsService.GetOne(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func CatsPost(c *gin.Context) {
	cat := new(models.Cat)
	if c.ShouldBind(cat) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request body",
		})
		return
	}

	if cat.ID == uuid.Nil {
		cat.ID = uuid.New()
	}

	id, err := catsService.Add(c.Request.Context(), cat)
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

func CatsPut(c *gin.Context) {
	cat := new(models.Cat)
	if c.ShouldBind(&cat) != nil {
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

	if err := catsService.Update(c.Request.Context(), id, cat); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "there was an error",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": fmt.Sprintf("id %s updated", id.String()),
	})
}
