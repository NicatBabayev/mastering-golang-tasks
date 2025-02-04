package jwt

import (
	"session-22/model"
	"session-22/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret, _ = utils.TimeToByte(time.Now())

func CreateToken(username string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 6)
	claims := &model.Claims{
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func Verify(token string) (string, error) {

	parsedToken := token[len("Authorization "):]

	claims, err := jwt.ParseWithClaims(parsedToken, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	return claims.Claims.(*model.Claims).UserName, nil
}
