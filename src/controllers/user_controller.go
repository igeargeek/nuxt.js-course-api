package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"github.com/icecreamhotz/movie-ticket/models"
	"github.com/icecreamhotz/movie-ticket/utils"
)

type UserHandler struct {
	Service   models.UserRepository
	Validator ut.Translator
}

func NewUserHandler(repository models.UserRepository, validator ut.Translator) UserHandler {
	return UserHandler{
		Service:   repository,
		Validator: validator,
	}
}

func (handler *UserHandler) RegisterUserPost(c *gin.Context) {
	var user models.User
	var err error
	if err = c.ShouldBind(&user); err != nil {
		errs := err.(validator.ValidationErrors)
		invalidFields := make([]map[string]string, 0)
		for _, e := range errs {
			errors := map[string]string{}
			errors[utils.ToSnakeCase(e.Field())] = e.Translate(handler.Validator)
			invalidFields = append(invalidFields, errors)
		}
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseErrorFields(invalidFields))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseMessage("Create Successfully."))
}
