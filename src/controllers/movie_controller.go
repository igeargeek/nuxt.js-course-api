package controllers

import (
	"net/http"

	"app/src/models"
	"app/src/utils"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

type MovieHandler struct {
	Service   models.MovieRepository
	Validator ut.Translator
}

func NewMovieHandler(repository models.MovieRepository, validator ut.Translator) MovieHandler {
	return MovieHandler{
		Service:   repository,
		Validator: validator,
	}
}

func (handler *MovieHandler) CreateMoviePost(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBind(&movie); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseErrorValidation(handler.Validator, err))
		return
	}
	id, err := handler.Service.Create(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseObject(gin.H{
		"message": "Created Successful",
		"id":      id,
	}))
}
