package controllers

import (
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MatkulController interface {
	InsertMatkul(c *gin.Context)
	UpdateMatkul(c *gin.Context)
	DeleteMatkul(c *gin.Context)
	FindMatkulByID(c *gin.Context)
	FindMatkulByName(c *gin.Context)
}

type MatkulControllerImplementation struct {
	MataKuliahService services.MataKuliahService
}

func NewMatkulController(matkulService services.MataKuliahService) MatkulController {
	return &MatkulControllerImplementation{MataKuliahService: matkulService}
}

func (l *MatkulControllerImplementation) InsertMatkul(c *gin.Context) {
	var matkul requests.InsertMatkulRequest
	err := c.ShouldBind(&matkul)
	if err != nil {
		errorList := helpers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := l.MataKuliahService.InsertMatkul(c.Request.Context(), &matkul)
	if errMsg != nil && !isSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Create Data Dosen",
			Data:       nil,
		})
		return
	}
}

func (l *MatkulControllerImplementation) UpdateMatkul(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	var matkul requests.UpdateMatkulRequest
	err = c.ShouldBind(&matkul)
	if err != nil {
		errorList := helpers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	matkul.ID = ID
	isSuccess, errMsg := l.MataKuliahService.UpdateMatkul(c.Request.Context(), &matkul)
	if errMsg != nil && !isSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Update Data Dosen",
			Data:       nil,
		})
		return
	}
}

func (l *MatkulControllerImplementation) DeleteMatkul(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := l.MataKuliahService.DeleteMatkulByID(c.Request.Context(), ID)
	if errMsg != nil && !isSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Hapus Data Dosen",
			Data:       nil,
		})
		return
	}

}

func (l *MatkulControllerImplementation) FindMatkulByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	matkul, isIDValid, errMsg := l.MataKuliahService.FindMatkulByID(c.Request.Context(), ID)
	if errMsg != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isIDValid {
		c.JSON(http.StatusOK, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Get Data Dosen",
			Data:       matkul,
		})
		return
	}

}

func (l *MatkulControllerImplementation) FindMatkulByName(c *gin.Context) {

	name := c.Query("name")
	matkuls, errMsg := l.MataKuliahService.FindMatkulByName(c.Request.Context(), name)
	if errMsg != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if matkuls != nil {
		c.JSON(http.StatusOK, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Get Data Matkul",
			Data:       matkuls,
		})
		return
	}
}
