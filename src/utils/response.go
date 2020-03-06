package utils

import (
	"time"
)

type NotFound struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type ServerError struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type Token struct {
	Success      bool      `json:"success"`
	Message      string    `json:"message"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type UnProcessEntity struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Errors  []map[string]string `json:"errors"`
}

type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ResponseObject(obj interface{}) interface{} {
	return obj
}

func ResponseMessage(message string) Message {
	return Message{
		Success: false,
		Message: message,
	}
}

func ReponseNotFound(message string) NotFound {
	return NotFound{
		Success: false,
		Message: message,
		Data:    []string{},
	}
}

func ResponseServerError(message string) ServerError {
	return ServerError{
		Success: false,
		Message: message,
		Data:    []string{},
	}
}

func ResponseToken(message, access_token, refresh_token string, expired_at time.Time) Token {
	return Token{
		Success:      true,
		Message:      message,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ExpiredAt:    expired_at,
	}
}

func ResponseErrorFields(errors []map[string]string) UnProcessEntity {
	return UnProcessEntity{
		Success: false,
		Message: "Invalid Fields",
		Errors:  errors,
	}
}
