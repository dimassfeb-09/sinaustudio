package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	responseError "github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/repository"
)

type UsersServiceImplementation struct {
	DB              *sql.DB
	UsersRepository repository.UsersRepository
}

func NewUserServiceImplementation(DB *sql.DB, UsersRepository repository.UsersRepository) UsersService {
	return &UsersServiceImplementation{
		DB:              DB,
		UsersRepository: UsersRepository,
	}
}

type UsersService interface {
	InsertDataUser(ctx context.Context, r *requests.UserInsertRequest) (bool, *responseError.ErrorMsg)
	UpdateDataUser(ctx context.Context, r *requests.UserUpdateRequest) (bool, *responseError.ErrorMsg)
	DeleteDataUser(ctx context.Context, UUID string) (bool, *responseError.ErrorMsg)
	FindUserByUUID(ctx context.Context, UUID string) (*domain.Users, *responseError.ErrorMsg)
	IsEmailRegistered(ctx context.Context, email string) (isRegistered bool, errMsg *responseError.ErrorMsg)
	IsNPMRegistered(ctx context.Context, npm string) (isRegistered bool, errMsg *responseError.ErrorMsg)
}

func (U *UsersServiceImplementation) InsertDataUser(ctx context.Context, r *requests.UserInsertRequest) (bool, *responseError.ErrorMsg) {
	tx, err := U.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isUUIDRegistered, _ := U.UsersRepository.FindUserByUUID(ctx, U.DB, r.UUID)
	if isUUIDRegistered {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "UUID sudah terdaftar")
	}

	_, isEmailRegistered, _ := U.UsersRepository.IsEmailRegistered(ctx, U.DB, r.Email)
	if isEmailRegistered {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "Email sudah digunakan")
	}

	_, isNPMRegistered, _ := U.UsersRepository.IsNPMRegistered(ctx, U.DB, r.NPM)
	if isNPMRegistered {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "NPM sudah digunakan")
	}

	user := &domain.Users{
		UUID:    r.UUID,
		Name:    r.Name,
		Email:   r.Email,
		Role:    r.Role,
		NPM:     r.NPM,
		ClassID: r.ClassID,
	}

	_, errMsg := U.UsersRepository.InsertDataUser(ctx, tx, user)
	if errMsg != nil {
		return false, errMsg
	}

	return true, nil
}

func (U *UsersServiceImplementation) UpdateDataUser(ctx context.Context, r *requests.UserUpdateRequest) (bool, *responseError.ErrorMsg) {
	tx, err := U.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	response, isEmailRegistered, _ := U.UsersRepository.IsEmailRegistered(ctx, U.DB, r.Email)
	if isEmailRegistered {
		if response.UUID == r.UUID {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_PREVIOUS_FIELD_NOT_ALLOWED, "Gunakan email yang baru.")
		} else {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "Email sudah digunakan.")
		}
	}

	response, isNPMRegistered, _ := U.UsersRepository.IsNPMRegistered(ctx, U.DB, r.NPM)
	if r.NPM != "" {
		if isNPMRegistered {
			if response.UUID == r.UUID {
				return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_PREVIOUS_FIELD_NOT_ALLOWED, "Gunakan NPM yang baru.")
			} else {
				return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "NPM sudah digunakan.")
			}
		}
	}

	user := &domain.Users{
		UUID:    r.UUID,
		Name:    r.Name,
		Email:   r.Email,
		NPM:     r.NPM,
		Role:    r.Role,
		ClassID: r.ClassID,
	}

	_, errMsg := U.UsersRepository.UpdateDataUser(ctx, tx, user)
	if errMsg != nil {
		return false, errMsg
	}

	return true, nil
}

func (U *UsersServiceImplementation) DeleteDataUser(ctx context.Context, UUID string) (bool, *responseError.ErrorMsg) {
	tx, err := U.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isUUIDRegistered, _ := U.UsersRepository.FindUserByUUID(ctx, U.DB, UUID)
	if !isUUIDRegistered {
		return false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "UUID tidak ditemukan.")
	}

	_, errMsg := U.UsersRepository.DeleteDataUser(ctx, tx, UUID)
	if errMsg != nil {
		return false, errMsg
	}

	return true, nil
}

func (U *UsersServiceImplementation) FindUserByUUID(ctx context.Context, UUID string) (*domain.Users, *responseError.ErrorMsg) {
	response, isUUIDRegistered, errMsg := U.UsersRepository.FindUserByUUID(ctx, U.DB, UUID)
	if errMsg != nil {
		return nil, errMsg
	}

	if !isUUIDRegistered {
		return nil, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data UUID tidak ditemukan.")
	}

	return response, nil
}

func (U *UsersServiceImplementation) IsEmailRegistered(ctx context.Context, email string) (isRegistered bool, errMsg *responseError.ErrorMsg) {
	_, isEmailRegistered, errMsg := U.UsersRepository.IsEmailRegistered(ctx, U.DB, email)
	if errMsg != nil {
		return false, errMsg
	}

	if !isEmailRegistered {
		return false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data Email tidak ditemukan.")
	}

	return true, nil
}

func (U *UsersServiceImplementation) IsNPMRegistered(ctx context.Context, npm string) (isRegistered bool, errMsg *responseError.ErrorMsg) {
	_, isNPMRegistered, errMsg := U.UsersRepository.IsNPMRegistered(ctx, U.DB, npm)
	if errMsg != nil {
		return false, errMsg
	}

	if !isNPMRegistered {
		return false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data Email tidak ditemukan.")
	}

	return true, nil
}
