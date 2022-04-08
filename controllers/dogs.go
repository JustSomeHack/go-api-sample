package controllers

import (
	"fmt"
	"net/http"

	"github.com/JustSomeHack/go-api-sample/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Deletes a dog by ID
// @Description deletes a dog
// @Produce  json
// @Param        dog_id    path      string     true  "Dog ID"
// @Success 200 {object} interface{}	"ok"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /dogs/{dog_id} [delete]
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

	c.JSON(http.StatusNoContent, gin.H{
		"deleted": id.String(),
	})
}

func DogsCount(c *gin.Context) {

}

// @Summary Gets all the dogs in the database
// @Description get a list of dogs
// @Produce  json
// @Success 200 {object} []models.Dog	"ok"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /dogs [get]
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

// @Summary Gets a dog by ID
// @Description get a dog
// @Produce  json
// @Param        dog_id    path      string     true  "Dog ID"
// @Success 200 {object} models.Dog	"ok"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /dogs/{dog_id} [get]
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

// @Summary Adds a dog
// @Description adds a dog
// @Accept   json
// @Produce  json
// @Param        message  body      models.Dog  true  "Dog"
// @Success      204   {string}  string  "answer"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /dogs [post]
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

// @Summary Updates a dog by ID
// @Description updates a dog
// @Accept   json
// @Produce  json
// @Param        dog_id    path      string     true  "Dog ID"
// @Param        message  body      models.Dog  true  "Dog"
// @Success      204   {string}  string  "answer"
// @Failure      400   {string}   string  "ok"
// @Failure      404   {string}   string  "ok"
// @Failure      500   {string}   string  "ok"
// @Router /dogs/{dog_id} [put]
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
