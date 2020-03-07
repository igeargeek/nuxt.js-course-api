package controllers

import (
	"fmt"
	"net/http"
	"strings"

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

	c.SetCookie("access_token", accessToken, utils.AccessTokenMinute, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, utils.RefreshTokenMinute, "/", "localhost", false, true)

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
		if err.Error() == utils.ErrRowExists.Error() {
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
	refreshToken := c.GetHeader("refresh_token")

	if refreshToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	refreshTokenCookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	if refreshTokenCookie != refreshToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	claims := &utils.Claims{}

	parseToken, err := jwt.ParseWithClaims(refreshTokenCookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisisasecretkeyrefreshtoken"), nil
	})

	if err != nil {
		fmt.Println(err)
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	if !parseToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	user, err := handler.Service.GetID(claims.ID)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	c.SetCookie("access_token", accessToken, utils.AccessTokenMinute, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, utils.RefreshTokenMinute, "/", "localhost", false, true)

	c.JSON(http.StatusOK, utils.ResponseToken("Refresh token success.",
		accessToken,
		refreshToken,
		expirationTimeAccessToken))
}

func (handler *UserHandler) PayloadTokenPost(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}
	token = strings.TrimSpace(splitToken[1])

	tokenCookie, err := c.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	if token != tokenCookie {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	claims := &utils.Claims{}

	parseToken, err := jwt.ParseWithClaims(tokenCookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisisasecretkeyaccesstoken"), nil
	})

	if err != nil {
		fmt.Println(err)
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	if !parseToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseObject(gin.H{
		"id":         claims.ID,
		"name":       claims.Name,
		"username":   claims.Username,
		"created_at": claims.CreatedAt,
		"updated_at": claims.UpdatedAt,
		"deleted_at": nil,
	}))
}
