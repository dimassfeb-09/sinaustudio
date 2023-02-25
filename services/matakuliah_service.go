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
)

type MataKuliahService interface {
	InsertMatkul(ctx context.Context, r *requests.InsertMatkulRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateMatkul(ctx context.Context, r *requests.UpdateMatkulRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteMatkulByID(ctx context.Context, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindMatkulByID(ctx context.Context, ID int) (response *response.MatkulResponse, isRegistered bool, errMsg *response.ErrorMsg)
	FindMatkulByName(ctx context.Context, name string) (matkuls []*response.MatkulResponse, errMsg *response.ErrorMsg)
}

type MataKuliahServiceImplementation struct {
	DB               *sql.DB
	MatkulRepository repository.MataKuliahRepository
	M                api.MicroServiceServer
}

func NewMataKuliahServiceImplementation(DB *sql.DB, M api.MicroServiceServer) MataKuliahService {
	return &MataKuliahServiceImplementation{DB: DB, MatkulRepository: M.MatkulRepository(), M: M}
}

func (m *MataKuliahServiceImplementation) InsertMatkul(ctx context.Context, r *requests.InsertMatkulRequest) (bool, *response.ErrorMsg) {
	tx, err := m.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	matkul := &domain.Matkul{
		Name:       r.Name,
		KodeMatkul: r.KodeMatkul,
	}
	isSuccess, errMsg := m.MatkulRepository.InsertMatkul(ctx, tx, matkul)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return isSuccess, nil
}

func (m *MataKuliahServiceImplementation) UpdateMatkul(ctx context.Context, r *requests.UpdateMatkulRequest) (bool, *response.ErrorMsg) {
	tx, err := m.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	matkul := &domain.Matkul{
		ID:         r.ID,
		Name:       r.Name,
		KodeMatkul: r.KodeMatkul,
	}
	isSuccess, errMsg := m.MatkulRepository.UpdateMatkul(ctx, tx, matkul)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return isSuccess, nil
}

func (m *MataKuliahServiceImplementation) DeleteMatkulByID(ctx context.Context, ID int) (bool, *response.ErrorMsg) {
	tx, err := m.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isRegistered, _ := m.MatkulRepository.FindMatkulByID(ctx, m.DB, ID)
	if !isRegistered {
		return false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Matkul ID tidak ditemukan")
	}

	isSuccess, errMsg := m.MatkulRepository.DeleteMatkulByID(ctx, tx, ID)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return isSuccess, nil
}

func (m *MataKuliahServiceImplementation) FindMatkulByID(ctx context.Context, ID int) (*response.MatkulResponse, bool, *response.ErrorMsg) {
	matkul, isRegistered, errMsg := m.MatkulRepository.FindMatkulByID(ctx, m.DB, ID)
	if errMsg != nil && !isRegistered {
		return nil, false, errMsg
	}

	matkulResponse := &response.MatkulResponse{
		ID:         matkul.ID,
		Name:       matkul.Name,
		KodeMatkul: matkul.KodeMatkul,
	}

	return matkulResponse, true, nil
}

func (m *MataKuliahServiceImplementation) FindMatkulByName(ctx context.Context, name string) ([]*response.MatkulResponse, *response.ErrorMsg) {
	responseName, errMsg := m.MatkulRepository.FindMatkulByName(ctx, m.DB, name)
	if errMsg != nil {
		return nil, errMsg

	}

	var matkuls []*response.MatkulResponse
	for i := 0; i < len(responseName); i++ {
		matkul := response.MatkulResponse{
			ID:         responseName[i].ID,
			Name:       responseName[i].Name,
			KodeMatkul: responseName[i].KodeMatkul,
		}
		matkuls = append(matkuls, &matkul)
	}

	return matkuls, nil
}
