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

type ClassRepository interface {
	InsertClass(ctx context.Context, tx *sql.Tx, class *domain.Class) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateClass(ctx context.Context, tx *sql.Tx, class *domain.Class) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteClassByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindClassByID(ctx context.Context, db *sql.DB, ID int) (*domain.Class, bool, *response.ErrorMsg)
	FindClassByName(ctx context.Context, db *sql.DB, name string) (*domain.Class, bool, *response.ErrorMsg)
}

type ClassRepositoryImplementation struct {
}

func NewClassRepositoryImplementation() ClassRepository {
	return &ClassRepositoryImplementation{}
}

func (c *ClassRepositoryImplementation) InsertClass(ctx context.Context, tx *sql.Tx, class *domain.Class) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "INSERT INTO class(name, kode_kelas) VALUES(?, ?)"
	_, err := tx.ExecContext(ctx, querySql, &class.Name, &class.KodeKelas)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (c *ClassRepositoryImplementation) UpdateClass(ctx context.Context, tx *sql.Tx, class *domain.Class) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "UPDATE class SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, &class.Name, &class.ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (c *ClassRepositoryImplementation) DeleteClassByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "DELETE FROM class WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	return true, nil
}

func (c *ClassRepositoryImplementation) FindClassByID(ctx context.Context, db *sql.DB, ID int) (*domain.Class, bool, *response.ErrorMsg) {
	querySql := "SELECT id, name, kode_kelas FROM class WHERE id = ?"
	row, err := db.QueryContext(ctx, querySql, ID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var class domain.Class
	if row.Next() {
		err := row.Scan(&class.ID, &class.Name, &class.KodeKelas)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_GET_DATA, err)
		} else {
			return &class, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data Class By ID tidak ditemukan.")
	}
}

func (c *ClassRepositoryImplementation) FindClassByName(ctx context.Context, db *sql.DB, name string) (*domain.Class, bool, *response.ErrorMsg) {
	querySql := "SELECT id, name FROM class WHERE name = ?"
	row, err := db.QueryContext(ctx, querySql, name)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer row.Close()

	var class domain.Class
	if row.Next() {
		err := row.Scan(&class.ID, &class.Name)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_GET_DATA, err)
		} else {
			return &class, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data Class By ID tidak ditemukan.")
	}
}
