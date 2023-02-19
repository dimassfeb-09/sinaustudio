package helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UsersInfoJWT struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	NPM     string `json:"npm"`
	ClassID int    `json:"class_id"`
}

type CustomJWTClaims struct {
	jwt.RegisteredClaims
	UsersInfo UsersInfoJWT `json:"users_info"`
}

func JWTGenereateToken(info UsersInfoJWT) (string, error) {

	expiredTime := time.Now().Add(time.Hour * 24)

	claims := CustomJWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "SINAUSTUDIO",
			Subject:   "Sinau Studios is",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UsersInfo: UsersInfoJWT{
			ID:      info.ID,
			Name:    info.Name,
			Email:   info.Email,
			NPM:     info.NPM,
			ClassID: info.ClassID,
		},
	}
	mySigningKey := []byte("7asd23&*^*($^&)**#$_hjagsd$#23496723")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	myToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return myToken, nil
}
