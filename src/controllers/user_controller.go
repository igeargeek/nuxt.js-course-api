package controllers

import (
	"fmt"
	"net/http"

	"app/src/constants"
	"app/src/models"
	"app/src/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"go.mongodb.org/mongo-driver/mongo"
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

func (handler *UserHandler) LoginUserPost(c *gin.Context) {
	var credential struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	if err := c.ShouldBind(&credential); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseErrorValidation(handler.Validator, err))
		return
	}

	user, err := handler.Service.FindByUsername(credential.Username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, utils.ResponseMessage("Username or password incorrect."))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	passwordMatch := utils.CheckPasswordHash(credential.Password, user.Password)
	if !passwordMatch {
		c.JSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	accessToken, refreshToken, expirationTimeAccessToken, err := utils.SetAccessTokenAndRefreshToken(
		user.ID,
		user.Name,
		user.Username,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	mAccessToken := &models.AccessToken{
		UserID: user.ID,
		Token:  accessToken,
	}
	err = handler.Service.CreateAccessToken(mAccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	mRefreshToken := &models.RefreshToken{
		UserID: user.ID,
		Token:  refreshToken,
	}
	err = handler.Service.CreateRefreshToken(mRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseToken("Login successful.",
		accessToken,
		refreshToken,
		expirationTimeAccessToken))
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
		if err.Error() == constants.ErrRowExists.Error() {
			c.JSON(http.StatusBadRequest, utils.ResponseMessage("Username is exists."))
			return
		} else {
			c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
			return
		}
	}

	c.JSON(http.StatusCreated, utils.ResponseObject(gin.H{
		"message": "Created Successful",
		"id":      id,
	}))
}

func (handler *UserHandler) RefreshTokenPost(c *gin.Context) {
	oldRefreshToken := c.PostForm("refresh_token")

	if oldRefreshToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	payload, err := handler.Service.GetRefreshToken(oldRefreshToken)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	claims := &utils.Claims{}

	parseToken, err := jwt.ParseWithClaims(payload.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return utils.RefreshKey, nil
	})

	if err != nil {
		fmt.Println("in err", err)
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	if !parseToken.Valid {
		fmt.Println("in parse", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	user, err := handler.Service.GetID(claims.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, utils.ResponseMessage("Username or password incorrect."))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	accessToken, refreshToken, expirationTimeAccessToken, err := utils.SetAccessTokenAndRefreshToken(
		user.ID,
		user.Name,
		user.Username,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	mAccessToken := &models.AccessToken{
		UserID: user.ID,
		Token:  accessToken,
	}
	err = handler.Service.CreateAccessToken(mAccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	mRefreshToken := &models.RefreshToken{
		UserID: user.ID,
		Token:  refreshToken,
	}
	err = handler.Service.CreateRefreshToken(mRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	err = handler.Service.RemoveRefreshToken(payload.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseToken("Refresh token success.",
		accessToken,
		refreshToken,
		expirationTimeAccessToken))
}

func (handler *UserHandler) PayloadTokenGet(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseObject(gin.H{
		"id":         c.MustGet("userId"),
		"name":       c.MustGet("Name"),
		"username":   c.MustGet("Username"),
		"created_at": c.MustGet("CreatedAt"),
		"updated_at": c.MustGet("UpdatedAt"),
		"deleted_at": nil,
	}))
}
