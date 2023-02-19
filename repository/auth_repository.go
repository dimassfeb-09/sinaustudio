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
	AuthLoginUser(ctx context.Context, db *sql.DB, email string) (isSuccess bool, result *domain.AuthLoginUser, errMsg *response.ErrorMsg)
}

type AuthRepositoryImplementation struct {
}

func NewAuthRepositoryImplementation() AuthRepository {
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

func (a *AuthRepositoryImplementation) AuthLoginUser(ctx context.Context, db *sql.DB, email string) (isSuccess bool, result *domain.AuthLoginUser, errMsg *response.ErrorMsg) {
	sqlQuery := "SELECT id, email, password FROM users WHERE email = ?"
	rows, err := db.QueryContext(ctx, sqlQuery, email)
	if err != nil {
		return false, nil, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer rows.Close()

	var user domain.AuthLoginUser
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return false, nil, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
		}
		return true, &user, nil
	} else {
		return false, nil, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Username atau Password salah.")
	}
}
