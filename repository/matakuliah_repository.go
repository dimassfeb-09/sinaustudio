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

type MataKuliahRepository interface {
	InsertMatkul(ctx context.Context, tx *sql.Tx, matkul *domain.Matkul) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateMatkul(ctx context.Context, tx *sql.Tx, matkul *domain.Matkul) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteMatkulByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindMatkulByID(ctx context.Context, db *sql.DB, ID int) (matkul *domain.Matkul, isRegistered bool, errMsg *response.ErrorMsg)
	FindMatkulByName(ctx context.Context, db *sql.DB, name string) (matkuls []*domain.Matkul, errMsg *response.ErrorMsg)
}

type MataKuliahRepositoryImplementation struct {
}

func NewMataKuliahRepositoryImplementation() MataKuliahRepository {
	return &MataKuliahRepositoryImplementation{}
}

func (m *MataKuliahRepositoryImplementation) InsertMatkul(ctx context.Context, tx *sql.Tx, matkul *domain.Matkul) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "INSERT INTO matakuliah(name, kode_matkul) VALUES(?, ?)"
	_, err := tx.ExecContext(ctx, querySql, &matkul.Name, &matkul.KodeMatkul)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, "Internal Server Error")
	}

	return true, nil
}

func (m *MataKuliahRepositoryImplementation) UpdateMatkul(ctx context.Context, tx *sql.Tx, matkul *domain.Matkul) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "UPDATE matakuliah SET name = ?, kode_matkul = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, &matkul.Name, &matkul.KodeMatkul, matkul.ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, "Internal Server Error")
	}

	return true, nil
}

func (m *MataKuliahRepositoryImplementation) DeleteMatkulByID(ctx context.Context, tx *sql.Tx, ID int) (isSuccess bool, errMsg *response.ErrorMsg) {
	querySql := "DELETE FROM matakuliah WHERE id  = ?"
	_, err := tx.ExecContext(ctx, querySql, ID)
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, "Internal Server Error")
	}

	return true, nil
}

func (m *MataKuliahRepositoryImplementation) FindMatkulByID(ctx context.Context, db *sql.DB, ID int) (*domain.Matkul, bool, *response.ErrorMsg) {
	querySql := "SELECT id, name, kode_matkul FROM matakuliah WHERE id = ?"
	row, err := db.QueryContext(ctx, querySql, ID)
	if err != nil {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, "Internal Server Error")
	}
	defer row.Close()

	var matkul domain.Matkul
	if row.Next() {
		err := row.Scan(&matkul.ID, &matkul.Name, &matkul.KodeMatkul)
		if err != nil {
			return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, err)
		} else {
			return &matkul, true, nil
		}
	} else {
		return nil, false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Matkul dengan ID tidak ditemukan.")
	}
}

func (m *MataKuliahRepositoryImplementation) FindMatkulByName(ctx context.Context, db *sql.DB, name string) ([]*domain.Matkul, *response.ErrorMsg) {
	rows, err := db.QueryContext(ctx, "SELECT id, name, kode_matkul FROM matakuliah WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer rows.Close()

	var matkuls []*domain.Matkul
	for rows.Next() {
		var matkul domain.Matkul
		err := rows.Scan(&matkul.ID, &matkul.Name, &matkul.KodeMatkul)
		if err != nil {
			return nil, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
		}
		matkuls = append(matkuls, &matkul)
	}

	return matkuls, nil
}
