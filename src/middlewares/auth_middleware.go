package middlewares

import (
	"app/src/models"
	"app/src/utils"

	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthMiddleware(model models.UserReporer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
			return
		}
		accessToken := strings.TrimSpace(splitToken[1])

		payload, err := model.GetAccessToken(accessToken)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
				return
			}
			c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
			return
		}

		parseToken, err := jwt.Parse(payload.Token, func(token *jwt.Token) (interface{}, error) {
			return utils.AccessKey, nil
		})

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
				return
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
			return
			}
		} else if !parseToken.Valid {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
			return
		}

		c.Next()
	}
}