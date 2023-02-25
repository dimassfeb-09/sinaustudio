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

type RoomController interface {
	InsertRoom(c *gin.Context)
	UpdateRoom(c *gin.Context)
	DeleteRoom(c *gin.Context)
	FindRoomByID(c *gin.Context)
}

type RoomControllerImplementation struct {
	RoomService services.RoomService
}

func NewRoomController(roomService services.RoomService) RoomController {
	return &RoomControllerImplementation{RoomService: roomService}
}

func (l *RoomControllerImplementation) InsertRoom(c *gin.Context) {

	var room requests.InsertRoomRequest
	err := c.ShouldBind(&room)
	if err != nil {
		errorList := helpers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := l.RoomService.InsertRoom(c.Request.Context(), &room)
	if errMsg != nil && !isSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Create Data Room",
			Data:       nil,
		})
		return
	}
}

func (l *RoomControllerImplementation) UpdateRoom(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	var room requests.UpdateRoomRequest
	err = c.ShouldBind(&room)
	if err != nil {
		errorList := helpers.ErrorValidateHandler(err)
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", errorList)
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	room.ID = ID
	isSuccess, errMsg := l.RoomService.UpdateRoom(c.Request.Context(), &room)
	if errMsg != nil && !isSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		helpers.ToWebResponse(c, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Update Data Room",
			Data:       nil,
		})
		return
	}
}

func (l *RoomControllerImplementation) DeleteRoom(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	isSuccess, errMsg := l.RoomService.DeleteRoomByID(c.Request.Context(), ID)
	if errMsg != nil && !isSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Hapus Data Room",
			Data:       nil,
		})
		return
	}

}

func (l *RoomControllerImplementation) FindRoomByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errMsg := helpers.ToErrorMsg(http.StatusBadRequest, "ERR_BAD_REQUEST_FIELD", "Invalid ID, fill with number ID")
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	room, isIDValid, errMsg := l.RoomService.FindRoomByID(c.Request.Context(), ID)
	if errMsg != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}

	if isIDValid {
		c.JSON(http.StatusOK, &response.SuccessResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Msg:        "Sukses Get Data Room",
			Data:       room,
		})
		return
	}
}
