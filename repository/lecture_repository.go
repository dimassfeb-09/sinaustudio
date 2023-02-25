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

type LectureRepository interface {
	InsertLecture(ctx context.Context, tx *sql.Tx, lecture *domain.Lecture) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateLecture(ctx context.Context, tx *sql.Tx, lecture *domain.Lecture) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteLectureByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindLectureByID(ctx context.Context, db *sql.DB, ID int) (lecture *domain.Lecture, isRegistered bool, errMsg *response.ErrorMsg)
	FindLectureByUserID(ctx context.Context, db *sql.DB, userID int) (lecture *domain.Lecture, isRegistered bool, errMsg *response.ErrorMsg)
	FindLectureByName(ctx context.Context, db *sql.DB, name string) (lecture *domain.Lecture, isRegistered bool, errMsg *response.ErrorMsg)
}

type LectureRepositoryImplementation struct {
}

func NewLectureRepositoryImplementation() LectureRepository {
	return &LectureRepositoryImplementation{}
}

func (l *LectureRepositoryImplementation) InsertLecture(ctx context.Context, tx *sql.Tx, lecture *domain.Lecture) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "INSERT INTO lecture(name, user_id) VALUES(?, ?)"
	_, err := tx.ExecContext(ctx, querySql, lecture.Name, lecture.UserID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (l *LectureRepositoryImplementation) UpdateLecture(ctx context.Context, tx *sql.Tx, lecture *domain.Lecture) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "UPDATE lecture SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, &lecture.Name, &lecture.ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (l *LectureRepositoryImplementation) DeleteLectureByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "DELETE FROM lecture WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (l *LectureRepositoryImplementation) FindLectureByID(ctx context.Context, db *sql.DB, ID int) (*domain.Lecture, bool, *response.ErrorMsg) {
	querySql := "SELECT id, name FROM lecture WHERE id = ?"
	row, err := db.QueryContext(ctx, querySql, ID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var lecture domain.Lecture
	if row.Next() {
		err := row.Scan(&lecture.ID, &lecture.Name)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, err)
		} else {
			return &lecture, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Dosen dengan ID tidak ditemukan.")
	}
}

func (l *LectureRepositoryImplementation) FindLectureByUserID(ctx context.Context, db *sql.DB, userID int) (*domain.Lecture, bool, *response.ErrorMsg) {
	querySql := "SELECT id, name, user_id FROM lecture WHERE user_id = ?"
	row, err := db.QueryContext(ctx, querySql, userID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var lecture domain.Lecture
	if row.Next() {
		err := row.Scan(&lecture.ID, &lecture.Name, &lecture.UserID)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, err)
		} else {
			return &lecture, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Dosen dengan ID tidak ditemukan.")
	}
}

func (l *LectureRepositoryImplementation) FindLectureByName(ctx context.Context, db *sql.DB, name string) (*domain.Lecture, bool, *response.ErrorMsg) {
	querySql := "SELECT id, name FROM lecture WHERE name = ?"
	row, err := db.QueryContext(ctx, querySql, name)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var lecture domain.Lecture
	if row.Next() {
		err := row.Scan(&lecture.ID, &lecture.Name)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, err)
		} else {
			return &lecture, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Dosen dengan Nama tidak ditemukan.")
	}
}
