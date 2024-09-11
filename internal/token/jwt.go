package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "kode"

func NewJwtToken(id int, login, mail string, init, expire time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":     id,
		"Login":  login,
		"Mail":   mail,
		"Init":   init,
		"Expire": expire,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractJWT(tokeString string) (jwt.MapClaims, error) {
	newToken, err := jwt.Parse(tokeString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token or method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("ExtractJWT", err)
	}

	if claims, ok := newToken.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, err
}

func ValidateTokenAndId(tokenString string) (int, error) {
	claims, err := ExtractJWT(tokenString)
	if err != nil {
		return 0, err
	}

	expireStr := claims["Expire"].(string)
	expire, err := time.Parse(time.RFC3339, expireStr)
	if err != nil {
		return 0, err
	}
	if expire.Before(time.Now()) {
		return 0, errors.New("token has expired")
	}

	if id, ok := claims["Id"].(float64); ok {
		return int(id), nil
	}

	return 0, errors.New("cookie is damadged")
}
