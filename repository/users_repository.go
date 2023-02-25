package repository

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
)

type UsersRepository interface {
	InsertDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteDataUser(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindUserByID(ctx context.Context, db *sql.DB, ID int) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg)
	FindUserByEmail(ctx context.Context, db *sql.DB, email string) (userResponse *domain.Users, errMsg *response.ErrorMsg)
	IsEmailRegistered(ctx context.Context, db *sql.DB, email string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg)
	ChangePasswordUser(ctx context.Context, tx *sql.Tx, newPass string, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindUserByClassID(ctx context.Context, db *sql.DB, classID int) (userResponse []domain.Users, isRegistered bool, errMsg *response.ErrorMsg)
}
type UsersRepositoryImplementations struct {
}

func NewUsersRepositoryImplementations() UsersRepository {
	return &UsersRepositoryImplementations{}
}

func (u *UsersRepositoryImplementations) InsertDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "INSERT INTO users(name, email, password, class_id, role) VALUES(?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, sqlQuery, &user.Name, &user.Email, &user.Password, &user.ClassID, &user.Role)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) UpdateDataUser(ctx context.Context, tx *sql.Tx, user *domain.Users) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "UPDATE users SET name = ?, email = ?, role = ?, class_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, &user.Name, &user.Email, &user.Role, &user.ClassID, &user.ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) DeleteDataUser(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, ID)
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

func (u *UsersRepositoryImplementations) FindUserByID(ctx context.Context, db *sql.DB, ID int) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg) {
	querySql := "SELECT id, name, email, password, class_id FROM users WHERE id = ?"
	row, err := db.QueryContext(ctx, querySql, ID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var user domain.Users
	if row.Next() {
		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.ClassID)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_GET_DATA, err)
		} else {
			return &user, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data User ID tidak ditemukan.")
	}
}

func (u *UsersRepositoryImplementations) FindUserByEmail(ctx context.Context, db *sql.DB, email string) (userResponse *domain.Users, errMsg *response.ErrorMsg) {
	querySql := "SELECT id, name, email, role, class_id FROM users WHERE email = ?"
	rows, err := db.QueryContext(ctx, querySql, email)
	if err != nil {
		return nil, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer rows.Close()

	var user domain.Users
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.ClassID)
		if err != nil {
			return nil, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
		}
		return &user, nil
	} else {
		return nil, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Email tidak ditemukan.")
	}

}

func (u *UsersRepositoryImplementations) IsEmailRegistered(ctx context.Context, db *sql.DB, email string) (userResponse *domain.Users, isRegistered bool, errMsg *response.ErrorMsg) {
	querySql := "SELECT id, email FROM users WHERE email = ?"
	row, err := db.QueryContext(ctx, querySql, email)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, "ERR_INTERNAL_SERVER_ERROR", err)
	}
	defer row.Close()

	var user domain.Users
	if row.Next() {
		err := row.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, "ERR_INTERNAL_SERVER_ERROR", err)
		}
		return &user, true, nil
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, "ERR_NOT_FOUND", "Data tidak ditemukan")
	}
}

func (u *UsersRepositoryImplementations) ChangePasswordUser(ctx context.Context, tx *sql.Tx, newPass string, ID int) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "UPDATE users SET password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, newPass, ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (u *UsersRepositoryImplementations) FindUserByClassID(ctx context.Context, db *sql.DB, classID int) (userResponse []domain.Users, isRegistered bool, errMsg *response.ErrorMsg) {
	querySql := "SELECT id, name, class_id FROM users WHERE class_id = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, querySql, classID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer rows.Close()

	var users []domain.Users
	for rows.Next() {
		var user domain.Users
		err := rows.Scan(&user.ID, &user.Name, &user.ClassID)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
		} else {
			users = append(users, user)
		}
	}
	return users, true, nil
}
