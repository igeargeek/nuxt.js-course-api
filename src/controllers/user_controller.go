package controllers

import (
	"net/http"

	"app/src/models"
	"app/src/utils"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

type UserHandler struct {
	Service   models.UserReporer
	Validator ut.Translator
}

func NewUserHandler(repository models.UserReporer, validator ut.Translator) UserHandler {
	return UserHandler{
		Service:   repository,
		Validator: validator,
	}
}

func (handler *UserHandler) RegisterUserPost(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseErrorValidation(handler.Validator, err))
		return
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	user.Password = hashPassword

	id, err := handler.Service.Create(&user)
	if err != nil {
		if err == utils.ErrRowExists {
			c.JSON(http.StatusBadRequest, utils.ResponseMessage("Username is exists."))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseObject(gin.H{
		"message": "Created Successful",
		"id":      id,
	}))
}

func (handler *UserHandler) ListOfUsersGet(c *gin.Context) {
	c.JSON(http.StatusCreated, utils.ResponseMessage("Create Successfully."))
}
