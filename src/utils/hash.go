package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var AccessKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
var RefreshKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

const AccessTokenMinute = 24 * 60
const RefreshTokenMinute = 24 * 7 * 60

type Claims struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Username  string             `json:"username"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	ID primitive.ObjectID
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(claim *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(AccessKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(claim *RefreshTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(RefreshKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func SetAccessTokenAndRefreshToken(id primitive.ObjectID, name, username string, createdAt, updatedAt time.Time) (string, string, time.Time, error) {
	expirationTimeAccessToken := time.Now().Add(time.Minute * time.Duration(AccessTokenMinute))
	expirationTimeRefreshToken := time.Now().Add(time.Minute * time.Duration(RefreshTokenMinute))

	accessTokenClaim := &Claims{
		ID:        id,
		Name:      name,
		Username:  username,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeAccessToken.Unix(),
		},
	}

	accessToken, err := GenerateToken(accessTokenClaim)
	if err != nil {
		return "", "", time.Time{}, err
	}

	refreshTokenClaims := &RefreshTokenClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeRefreshToken.Unix(),
		},
	}

	refreshToken, err := GenerateRefreshToken(refreshTokenClaims)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken, refreshToken, expirationTimeAccessToken, nil
}

func GetUserPayload(token string) (*Claims, *jwt.Token, error) {
	// example use
	// token, ok := utils.SplitTokenFromHeader(tokenHeader)
	// if !ok {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
	// 	return
	// }

	// claims, parseToken, err := utils.GetUserPayload(token)
	// เช็คว่า token ที่ส่งมาถูกต้องหรือป่าว parse token และเช็ค
	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
	// 		return
	// 	}
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
	// 	return
	// }

	// end of handler error
	// เช็คว่า token ทั้งหมด header payload signature ถูกทั้งหมดหรือป่าว
	// if !parseToken.Valid {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage("Unauthorized."))
	// 	return
	// }
	// claims contain
	// end of example use

	claims := &Claims{}

	parseToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return AccessKey, nil
	})

	if err != nil || !parseToken.Valid {
		return nil, parseToken, err
	}

	return claims, parseToken, nil
}
