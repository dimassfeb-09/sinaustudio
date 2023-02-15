package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
)

type UsersRepository interface {
	InsertDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteDataUser(ctx context.Context, tx *sql.Tx, UUID string) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateEmailUser(ctx context.Context, tx *sql.Tx, recentEmail string, newEmail string) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateNameUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg)
	FindUserByUUID(ctx context.Context, db *sql.DB, UUID string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg)
	IsEmailRegistered(ctx context.Context, db *sql.DB, email string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg)
	IsNPMRegistered(ctx context.Context, db *sql.DB, npm string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg)
}

type UsersRepositoryImplementations struct {
}

func NewUsersRepositoryImplementations() UsersRepository {
	return &UsersRepositoryImplementations{}
}

func (u *UsersRepositoryImplementations) InsertDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "INSERT INTO users(uuid, name, email, class_id, npm, role) VALUES(?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, sqlQuery, &user.UUID, &user.Name, &user.Email, &user.ClassID, &user.NPM, &user.Role)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) UpdateDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "UPDATE users SET name = ?, email = ?, class_id = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, &user.Name, &user.Email, &user.ClassID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) DeleteDataUser(ctx context.Context, tx *sql.Tx, UUID string) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "DELETE FROM users WHERE uuid = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, UUID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) UpdateEmailUser(ctx context.Context, tx *sql.Tx, recentEmail string, newEmail string) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "UPDATE users SET email = ? WHERE email = ?"
	_, err := tx.ExecContext(ctx, querySql, newEmail, recentEmail)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) FindUserByUUID(ctx context.Context, db *sql.DB, UUID string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg) {
	querySql := "SELECT id, uuid, name, email, class_id FROM users WHERE uuid = ?"
	row, err := db.QueryContext(ctx, querySql, UUID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var user domain.Users
	if row.Next() {
		err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.ClassID)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_GET_DATA, err)
		} else {
			return &user, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_GET_DATA, errors.New("Data User UUID tidak ditemukan."))
	}
}

func (u *UsersRepositoryImplementations) UpdateNameUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "UPDATE users SET name = ? WHERE uuid = ?"
	_, err := tx.QueryContext(ctx, querySql, &user.Name, &user.UUID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) IsEmailRegistered(ctx context.Context, db *sql.DB, email string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg) {
	querySql := "SELECT uuid, email FROM users WHERE email = ?"
	row, err := db.QueryContext(ctx, querySql, email)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, "ERR_INTERNAL_SERVER_ERROR", err)
	}
	defer row.Close()

	var user domain.Users
	if row.Next() {
		err := row.Scan(&user.UUID, &user.Email)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, "ERR_INTERNAL_SERVER_ERROR", err)
		}
		return &user, true, nil
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, "ERR_NOT_FOUND", "Data tidak ditemukan")
	}
}

func (u *UsersRepositoryImplementations) IsNPMRegistered(ctx context.Context, db *sql.DB, npm string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg) {
	querySql := "SELECT uuid, npm FROM users WHERE npm = ?"
	row, err := db.QueryContext(ctx, querySql, npm)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err)
	}
	defer row.Close()

	var user domain.Users
	if row.Next() {
		err := row.Scan(&user.UUID, &user.NPM)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, "ERR_INTERNAL_SERVER_ERROR", err)
		}
		return &user, true, nil
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, "ERR_NOT_FOUND", "Data tidak ditemukan")
	}
}
