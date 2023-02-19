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

type AuthRepository interface {
	AuthRegisterUser(ctx context.Context, tx *sql.Tx, user *domain.AuthRegisterUser) (isSuccess bool, errMsg *response.ErrorMsg)
	AuthLoginUser(ctx context.Context, db *sql.DB, email string, password string) (isSuccess bool, errMsg *response.ErrorMsg)
}

type AuthRepositoryImplementation struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImplementation{}
}

func (a *AuthRepositoryImplementation) AuthRegisterUser(ctx context.Context, tx *sql.Tx, user *domain.AuthRegisterUser) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "INSERT INTO users(name, email, password, class_id, npm, role) VALUES (?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, sqlQuery, &user.Name, &user.Email, &user.Password, &user.ClassID, &user.NPM, &user.Role)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (a *AuthRepositoryImplementation) AuthLoginUser(ctx context.Context, db *sql.DB, email string, password string) (isSuccess bool, errMsg *response.ErrorMsg) {
	sqlQuery := "SELECT id FROM users WHERE email = ? AND password = ?"
	rows, err := db.QueryContext(ctx, sqlQuery, email, password)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	} else {
		return false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data tidak ditemukan")
	}
}
