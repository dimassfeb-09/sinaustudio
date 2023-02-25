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

type ClassService interface {
	AddClass(ctx context.Context, r *requests.InsertClassRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateClass(ctx context.Context, r *requests.UpdateClassRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteClassByID(ctx context.Context, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindClassByID(ctx context.Context, ID int) (*domain.Class, bool, *response.ErrorMsg)
	FindClassByName(ctx context.Context, name string) (*domain.Class, bool, *response.ErrorMsg)
}

type ClassServiceImplementation struct {
	DB              *sql.DB
	ClassRepository repository.ClassRepository
	M               api.MicroServiceServer
}

func NewClassServiceImplementation(DB *sql.DB, M api.MicroServiceServer) ClassService {
	return &ClassServiceImplementation{DB: DB, ClassRepository: M.ClassRepository(), M: M}
}

func (c *ClassServiceImplementation) AddClass(ctx context.Context, r *requests.InsertClassRequest) (bool, *response.ErrorMsg) {
	tx, err := c.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	r.KodeKelas = helpers.RandStr(10)
	class := &domain.Class{
		Name:      r.Name,
		KodeKelas: r.KodeKelas,
	}
	isSuccess, errMsg := c.ClassRepository.InsertClass(ctx, tx, class)
	if !isSuccess && errMsg != nil {
		return false, errMsg
	}

	return true, nil
}

func (c *ClassServiceImplementation) UpdateClass(ctx context.Context, r *requests.UpdateClassRequest) (bool, *response.ErrorMsg) {
	tx, err := c.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isIDRegistered, errMsg := c.ClassRepository.FindClassByID(ctx, c.DB, r.ID)
	if !isIDRegistered {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_NOT_FOUND, "Class ID tidak ditemukan.")
	}

	class := &domain.Class{
		ID:   r.ID,
		Name: r.Name,
	}

	isSuccess, errMsg := c.ClassRepository.UpdateClass(ctx, tx, class)
	if !isSuccess && errMsg != nil {
		return false, errMsg
	}

	return true, nil
}

func (c *ClassServiceImplementation) DeleteClassByID(ctx context.Context, ID int) (bool, *response.ErrorMsg) {
	tx, err := c.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isClassIDAlreadyUse, _ := c.M.UserRepository().FindUserByClassID(ctx, c.DB, ID)
	if isClassIDAlreadyUse {
		return false, helpers.ToErrorMsg(http.StatusConflict, exception.ERR_CONFLICT, "Kelas tidak dapat dihapus, karena berelasi dengan data user.")
	}

	_, isIDRegistered, _ := c.ClassRepository.FindClassByID(ctx, c.DB, ID)
	if !isIDRegistered {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_NOT_FOUND, "Class by ID tidak ditemukan.")
	}

	isSuccess, errMsg := c.ClassRepository.DeleteClassByID(ctx, tx, ID)
	if !isSuccess && errMsg != nil {
		return false, errMsg
	}

	return true, nil
}

func (c *ClassServiceImplementation) FindClassByID(ctx context.Context, ID int) (*domain.Class, bool, *response.ErrorMsg) {
	r, isIDValid, _ := c.ClassRepository.FindClassByID(ctx, c.DB, ID)
	if !isIDValid {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_NOT_FOUND, "Kelas by ID tidak ditemukan.")
	}
	return r, true, nil
}

func (c *ClassServiceImplementation) FindClassByName(ctx context.Context, name string) (*domain.Class, bool, *response.ErrorMsg) {
	r, isNameValid, _ := c.ClassRepository.FindClassByName(ctx, c.DB, name)
	if !isNameValid {
		return nil, false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_NOT_FOUND, "Kelas by Name tidak ditemukan.")
	}
	return r, true, nil
}
