package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var accessKey = []byte("thisisasecretkeyaccesstoken")
var refreshKey = []byte("thisisasecretkeyrefreshtoken")

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
	tokenString, err := token.SignedString(accessKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(claim *RefreshTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(refreshKey)
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
