package controllers

import (
	"database/sql"
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/handlers"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUsersControllerImplementation(usersService services.UsersService) UsersController {
	return &UsersControllerImplementation{UsersService: usersService}
}

type UsersControllerImplementation struct {
	UsersService services.UsersService
	DB           *sql.DB
}

type UsersController interface {
	InsertDataUser(c *gin.Context)
	UpdateDataUser(c *gin.Context)
	DeleteDataUser(c *gin.Context)
	UpdateEmailUser(c *gin.Context)
	UpdateNameUser(c *gin.Context)
	UpdateNPMUser(c *gin.Context)
	FindUserByUUID(c *gin.Context)
	IsEmailRegistered(c *gin.Context)
	IsNPMRegistered(c *gin.Context)
}

func (u *UsersControllerImplementation) InsertDataUser(c *gin.Context) {
	var user requests.UserInsertRequest
	err := c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := u.UsersService.InsertDataUser(c.Request.Context(), &user)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: 200,
			Msg:        "Sukses Create Data User",
			Data:       nil,
		})
	}
}

func (u *UsersControllerImplementation) UpdateDataUser(c *gin.Context) {
	var user requests.UserUpdateRequest
	err := c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := u.UsersService.UpdateDataUser(c.Request.Context(), &user)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: 200,
			Msg:        "Sukses Update Data User",
			Data:       nil,
		})
	}
}

func (u *UsersControllerImplementation) DeleteDataUser(c *gin.Context) {
	var user requests.UserDeleteRequest
	err := c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := u.UsersService.DeleteDataUser(c.Request.Context(), user.UUID)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: 200,
			Msg:        "Sukses Delete Data User",
			Data:       nil,
		})
	}
}

func (u *UsersControllerImplementation) UpdateEmailUser(c *gin.Context) {
	var user requests.UserUpdateEmailRequest
	err := c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := u.UsersService.UpdateEmailUser(c.Request.Context(), user.Email, user.UUID)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: 200,
			Msg:        "Sukses Update Email User",
			Data:       nil,
		})
	}
}

func (u *UsersControllerImplementation) UpdateNameUser(c *gin.Context) {
	var user requests.UserUpdateNameRequest
	err := c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := u.UsersService.UpdateNameUser(c.Request.Context(), user.Name, user.UUID)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: 200,
			Msg:        "Sukses Update Name User",
			Data:       nil,
		})
	}
}

func (u *UsersControllerImplementation) UpdateNPMUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UsersControllerImplementation) FindUserByUUID(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UsersControllerImplementation) IsEmailRegistered(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UsersControllerImplementation) IsNPMRegistered(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
