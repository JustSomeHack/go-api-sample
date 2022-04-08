package controllers

import (
	"fmt"
	"net/http"

	"github.com/JustSomeHack/go-api-sample/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Deletes a cat by ID
// @Description deletes a cat
// @Produce  json
// @Success 200 {object} interface{}	"ok"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /cats/{cat_id} [delete]
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

// @Summary Gets all the cats in the database
// @Description get a list of cats
// @Produce  json
// @Success 200 {object} []models.Cat	"ok"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /cats [get]
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

// @Summary Gets a cat by ID
// @Description get a cat
// @Produce  json
// @Param        cat_id    path      string     true  "Cat ID"
// @Success 200 {object} models.Cat	"ok"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /cats/{cat_id} [get]
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

// @Summary Adds a cat
// @Description adds a cat
// @Accept   json
// @Produce  json
// @Param        message  body      models.Cat  true  "Cat"
// @Success      204   {string}  string  "answer"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /cats [post]
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

// @Summary Updates a cat by ID
// @Description updates a cat
// @Accept   json
// @Produce  json
// @Param        cat_id    path      string     true  "Cat ID"
// @Param        message  body      models.Cat  true  "Cat"
// @Success      204   {string}  string  "answer"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /cats/{cat_id} [put]
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
