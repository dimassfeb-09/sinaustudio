package controllers

import (
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/handlers"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ClassController interface {
	AddClass(c *gin.Context)
	UpdateClass(c *gin.Context)
	DeleteClassByID(c *gin.Context)
	FindClassByID(c *gin.Context)
	FindClassByName(c *gin.Context)
}

type ClassControllerImplementation struct {
	ClassService services.ClassService
}

func NewClassControllerImplementation(classService services.ClassService) ClassController {
	return &ClassControllerImplementation{ClassService: classService}
}

func (k ClassControllerImplementation) AddClass(c *gin.Context) {

	var class requests.InsertClassRequest
	err := c.ShouldBind(&class)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := k.ClassService.AddClass(c.Request.Context(), class.Name)
	if !isSuccess && errMsg != nil {
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Create Data Kelas",
			Data:       nil,
		})
		return
	}

}

func (k ClassControllerImplementation) UpdateClass(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, must integer")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	var class requests.UpdateClassRequest
	err = c.ShouldBind(&class)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	class.ID = ID
	isSuccess, errMsg := k.ClassService.UpdateClass(c.Request.Context(), &class)
	if !isSuccess && errMsg != nil {
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Update Data Kelas",
			Data:       nil,
		})
	}
}

func (k ClassControllerImplementation) DeleteClassByID(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Invalid ID, fill with number/integer")
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	isSuccess, errMsg := k.ClassService.DeleteClassByID(c.Request.Context(), ID)
	if !isSuccess && errMsg != nil {
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Delete Data Kelas",
			Data:       nil,
		})
	}

}

func (k ClassControllerImplementation) FindClassByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Invalid ID, fill with number/integer")
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	result, isIDRegistered, errMsg := k.ClassService.FindClassByID(c.Request.Context(), ID)
	if !isIDRegistered && errMsg != nil {
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	if isIDRegistered {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Get Data Kelas",
			Data:       result,
		})
		return
	}

	return
}

func (k ClassControllerImplementation) FindClassByName(c *gin.Context) {
	name := c.Query("name")
	result, isIDRegistered, errMsg := k.ClassService.FindClassByName(c.Request.Context(), name)
	if !isIDRegistered && errMsg != nil {
		c.AbortWithStatusJSON(errMsg.StatusCode, errMsg)
		return
	}

	if isIDRegistered {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Get Data Kelas",
			Data:       result,
		})
		return
	}
}
