package api

import (
	"errors"
	"fmt"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type CustomJWTClaims struct {
	*jwt.RegisteredClaims
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	ClassID int    `json:"class_id"`
}

type UserInfo struct {
	ID      int
	Name    string
	Email   string
	Role    string
	ClassID int
}

var mySigningKey = []byte("7asd23&*^*($^&)**#$_hjagsd$#23496723")

func CheckingJWTToken(tokenBearer string, c *gin.Context) {

	tokenResult := strings.TrimLeft(tokenBearer, " ")
	token, err := jwt.Parse(tokenResult, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		c.Header("Authorization", fmt.Sprintf("Bearer %s", token.Raw))
		c.SetCookie("Authorization", fmt.Sprintf("Bearer %s", token.Raw), 10, "/", "localhost:8080", true, true)

		return mySigningKey, nil
	})

	if err != nil {
		if errors.Is(err, &jwt.ValidationError{}) {
			errMsg := helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, err.(*jwt.ValidationError).Error())
			c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
			return
		}
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Token tidak valid!")
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}
}

func JWTGenereateToken(info *UserInfo) (string, error) {

	claims := &CustomJWTClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "SINAUSTUDIO",
			Subject:   "Sinau Studios is",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		ID:      info.ID,
		Name:    info.Name,
		Email:   info.Name,
		Role:    info.Role,
		ClassID: info.ClassID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	myToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return myToken, nil
}
