package middleware

import (
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func MiddlewareAuthorizationfunc(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		errMsg := helpers.ToErrorMsg(http.StatusUnauthorized, exception.ERR_UNAUTHORIZED_BEARER, "Key: Header with key Authorization: Bearer Token, Tag: Required")
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}
	bearers := strings.Split(authorization, "Bearer")
	if bearers[1] == "" {
		errMsg := helpers.ToErrorMsg(http.StatusUnauthorized, exception.ERR_UNAUTHORIZED_BEARER, "Key: Token not found, Tag: Required")
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	} else {
		CheckingJWTToken(bearers[1], c)
		c.Next()
	}
}
