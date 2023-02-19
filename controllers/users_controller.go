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
	"strconv"
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
	FindUserByID(c *gin.Context)
	IsEmailRegistered(c *gin.Context)
	IsNPMRegistered(c *gin.Context)
	ChangePasswordUser(c *gin.Context)
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
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
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
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Update Data User",
			Data:       nil,
		})
	}
}

func (u *UsersControllerImplementation) DeleteDataUser(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	var user requests.UserDeleteRequest
	err = c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := u.UsersService.DeleteDataUser(c.Request.Context(), user.ConfirmPassword, ID)
	if errMsg != nil {
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Delete Data User",
			Data:       nil,
		})
	}
}

func (u *UsersControllerImplementation) FindUserByID(c *gin.Context) {
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

func (u *UsersControllerImplementation) ChangePasswordUser(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	var user requests.UserChangePassword
	err = c.ShouldBind(&user)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := u.UsersService.ChangePasswordUser(c.Request.Context(), ID, user.RecentPassword, user.NewPassword)
	if !isSuccess && errMsg != nil {
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)

		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Change Password Data User",
			Data:       nil,
		})
	}

}
