package controllers

import (
	"fmt"
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/handlers"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LectureController interface {
	InsertLecture(c *gin.Context)
	UpdateLecture(c *gin.Context)
	DeleteLecture(c *gin.Context)
	FindLectureByID(c *gin.Context)
	FindLectureByName(c *gin.Context)
}

type LectureControllerImplementation struct {
	LectureService services.LectureService
}

func NewLectureController(lectureService services.LectureService) LectureController {
	return &LectureControllerImplementation{LectureService: lectureService}
}

func (l *LectureControllerImplementation) InsertLecture(c *gin.Context) {

	var lecture requests.InsertLectureRequest
	err := c.ShouldBind(&lecture)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := l.LectureService.InsertLecture(c.Request.Context(), &lecture)
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

func (l *LectureControllerImplementation) UpdateLecture(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	var lecture requests.UpdateLectureRequest
	err = c.ShouldBind(&lecture)
	if err != nil {
		errorList := handlers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	lecture.ID = ID
	isSuccess, errMsg := l.LectureService.UpdateLecture(c.Request.Context(), &lecture)
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

func (l *LectureControllerImplementation) DeleteLecture(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := l.LectureService.DeleteLectureByID(c.Request.Context(), ID)
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

func (l *LectureControllerImplementation) FindLectureByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	lecture, isIDValid, errMsg := l.LectureService.FindLectureByID(c.Request.Context(), ID)
	if errMsg != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	fmt.Println(lecture)
	fmt.Println("NOT FOUND")

	if isIDValid {
		c.JSON(http.StatusOK, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Get Data Dosen",
			Data:       lecture,
		})
		return
	}

}

func (l *LectureControllerImplementation) FindLectureByName(c *gin.Context) {

	name := c.Query("name")
	lecture, isIDValid, errMsg := l.LectureService.FindLectureByName(c.Request.Context(), name)
	if errMsg != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isIDValid {
		c.JSON(http.StatusOK, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Get Data Dosen",
			Data:       lecture,
		})
		return
	}
}
