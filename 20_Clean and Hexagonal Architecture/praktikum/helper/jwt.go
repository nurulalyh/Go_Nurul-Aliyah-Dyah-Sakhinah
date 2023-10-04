package helper

import (
	"clean_arc/config"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func GenerateJWT(id string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resToken, err := rawToken.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		logrus.Println("Generate jwt error, ", err.Error())
		return ""
	}

	return resToken
}

func DecodeJWT(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		user_id := data["id"].(string)

		return user_id
	}

	return ""
}