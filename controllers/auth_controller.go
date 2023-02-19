package controllers

import (
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/handlers"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/middleware"
	"github.com/dimassfeb-09/sinaustudio.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	AuthRegisterUser(c *gin.Context)
	AuthLoginUser(c *gin.Context)
}

type AuthControllerImplementation struct {
	AuthService services.AuthService
}

func NewAuthControllerImplementation(authService services.AuthService) AuthController {
	return &AuthControllerImplementation{AuthService: authService}
}

func (a *AuthControllerImplementation) AuthRegisterUser(c *gin.Context) {
	var auth requests.AuthRegisterRequest
	err := c.ShouldBind(&auth)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := a.AuthService.AuthRegisterUser(c.Request.Context(), &auth)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Create Data User",
			Data:       nil,
		})
	}
}

func (a *AuthControllerImplementation) AuthLoginUser(c *gin.Context) {
	var user requests.AuthLoginRequest
	err := c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	userInfo, errMsg := a.AuthService.AuthLoginUser(c.Request.Context(), user.Email, user.Password)
	if errMsg != nil && userInfo == nil {

		c.JSON(errMsg.StatusCode, errMsg)
		return
	}

	if userInfo != nil {

		userInfoJWT := &middleware.UserInfo{
			ID:      userInfo.ID,
			Name:    userInfo.Name,
			Email:   userInfo.Email,
			NPM:     userInfo.NPM,
			Role:    userInfo.Role,
			ClassID: userInfo.ClassID,
		}
		token, err := middleware.JWTGenereateToken(userInfoJWT)
		if err != nil {
			c.JSON(errMsg.StatusCode, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, err))
			return
		}

		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Login",
			Data: map[string]any{
				"token": token,
			},
		})
		return
	}
}
