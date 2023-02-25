package services

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/sinaustudio.git/api"
	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/repository"
	"net/http"
	"time"
)

type RoomService interface {
	InsertRoom(ctx context.Context, r *requests.InsertRoomRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateRoom(ctx context.Context, r *requests.UpdateRoomRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteRoomByID(ctx context.Context, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindRoomByID(ctx context.Context, ID int) (r *response.RoomResponse, isValid bool, errMsg *response.ErrorMsg)
}

type RoomServiceImplementation struct {
	DB             *sql.DB
	RoomRepository repository.RoomRepository
	M              api.MicroServiceServer
}

func NewRoomServiceImplementation(DB *sql.DB, m api.MicroServiceServer) RoomService {
	return &RoomServiceImplementation{DB: DB, RoomRepository: m.RoomRepository(), M: m}
}

func (l *RoomServiceImplementation) InsertRoom(ctx context.Context, r *requests.InsertRoomRequest) (bool, *response.ErrorMsg) {
	tx, err := l.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	startRoom, err := time.Parse("2006-01-02 15:04:05", r.StartRoom)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Format waktu tidak sesuai, gunakan: yyyy-mm-dd HH:mm:ss")
	}
	endRoom, err := time.Parse("2006-01-02 15:04:05", r.EndRoom)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Format waktu tidak sesuai, gunakan: yyyy-mm-dd HH:mm:ss")
	}

	r.StartRoom = startRoom.Format("2006-01-02 15:04:05")
	r.EndRoom = endRoom.Format("2006-01-02 15:04:05")

	room := &domain.Room{
		ID:        r.ID,
		Name:      r.Name,
		URL:       r.URL,
		LectureID: r.LectureID,
		StartRoom: r.StartRoom,
		EndRoom:   r.EndRoom,
	}

	isSuccess, errMsg := l.RoomRepository.InsertRoom(ctx, tx, room)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return true, nil
}

func (l *RoomServiceImplementation) UpdateRoom(ctx context.Context, r *requests.UpdateRoomRequest) (bool, *response.ErrorMsg) {
	tx, err := l.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isIDValid, errMsg := l.RoomRepository.FindRoomByID(ctx, l.DB, r.ID)
	if errMsg != nil && !isIDValid {
		return false, errMsg
	}

	room := &domain.Room{
		ID:        r.ID,
		Name:      r.Name,
		URL:       r.URL,
		LectureID: r.LectureID,
		StartRoom: r.StartRoom,
		EndRoom:   r.EndRoom,
	}

	isSuccess, errMsg := l.RoomRepository.UpdateRoom(ctx, tx, room)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return true, nil
}

func (l *RoomServiceImplementation) DeleteRoomByID(ctx context.Context, ID int) (bool, *response.ErrorMsg) {
	tx, err := l.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	isSuccess, errMsg := l.RoomRepository.DeleteRoomByID(ctx, tx, ID)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return true, nil
}

func (l *RoomServiceImplementation) FindRoomByID(ctx context.Context, ID int) (r *response.RoomResponse, isValid bool, errMsg *response.ErrorMsg) {
	room, isIDValid, errMsg := l.RoomRepository.FindRoomByID(ctx, l.DB, ID)
	if isIDValid {
		roomResponse := &response.RoomResponse{
			ID:        room.ID,
			Name:      room.Name,
			URL:       room.URL,
			LectureID: room.LectureID,
			StartRoom: room.StartRoom,
			EndRoom:   room.EndRoom,
		}
		return roomResponse, true, nil
	} else {
		return nil, false, errMsg
	}
}
