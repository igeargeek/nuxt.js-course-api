package utils

import (
	"strings"
	"time"
)

func SplitTokenFromHeader(token string) (string, bool) {
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return "", false
	}
	token = strings.TrimSpace(splitToken[1])

	return token, true
}

func GetTimeNowFormatYYYYMMDDHHIIMM() string {
	t := time.Now()
	s := t.Format("20060102150405")
	return s
}

func CheckInArrayString(arr []string, keyword string) bool {
	for _, a := range arr {
		if a == keyword {
			return true
		}
	}
	return false
}
