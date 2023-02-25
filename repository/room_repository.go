package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"net/http"
)

type RoomRepository interface {
	InsertRoom(ctx context.Context, tx *sql.Tx, room *domain.Room) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateRoom(ctx context.Context, tx *sql.Tx, room *domain.Room) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteRoomByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindRoomByID(ctx context.Context, db *sql.DB, ID int) (room *domain.Room, isRegistered bool, errMsg *response.ErrorMsg)
}

type RoomRepositoryImplementation struct {
}

func NewRoomRepositoryImplementation() RoomRepository {
	return &RoomRepositoryImplementation{}
}

func (r *RoomRepositoryImplementation) InsertRoom(ctx context.Context, tx *sql.Tx, room *domain.Room) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "INSERT INTO room(name, url, lecture_id, start_room, end_room) VALUES(?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, querySql, &room.Name, &room.URL, &room.LectureID, &room.StartRoom, &room.EndRoom)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (r *RoomRepositoryImplementation) UpdateRoom(ctx context.Context, tx *sql.Tx, room *domain.Room) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "UPDATE room SET name = ?, lecture_id = ?, start_room = ?, end_room = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, &room.Name, &room.LectureID, &room.StartRoom, &room.EndRoom, &room.ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (r *RoomRepositoryImplementation) DeleteRoomByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "DELETE FROM room WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (r *RoomRepositoryImplementation) FindRoomByID(ctx context.Context, db *sql.DB, ID int) (Room *domain.Room, isRegistered bool, errMsg *response.ErrorMsg) {
	querySql := "SELECT id, name, url, lecture_id, start_room, end_room FROM room WHERE id = ?"
	row, err := db.QueryContext(ctx, querySql, ID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var room domain.Room
	if row.Next() {
		err := row.Scan(&room.ID, &room.Name, &room.URL, &room.LectureID, &room.StartRoom, &room.EndRoom)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, err)
		} else {
			return &room, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Room dengan ID tidak ditemukan.")
	}
}
