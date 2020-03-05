package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icecreamhotz/movie-ticket/utils"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (handler *ServiceHandler) RegisterUserPost(c *gin.Context) {
	var user User
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
