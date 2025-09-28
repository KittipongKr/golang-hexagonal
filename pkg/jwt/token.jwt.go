package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignTokenHMAC(secret string, payload map[string]interface{}, exp int64) (string, error) {
	payload["exp"] = time.Now().Add(time.Minute * time.Duration(exp)).Unix()
	payload["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	return token.SignedString([]byte(secret))
}
