package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/repository"
)

type UsersService interface {
	InsertDataUser(ctx context.Context, r *requests.UserInsertRequest) (bool, *response.ErrorMsg)
	UpdateDataUser(ctx context.Context, r *requests.UserUpdateRequest) (bool, *response.ErrorMsg)
	// DeleteDataUser(ctx context.Context, UUID string) (bool, *models.ErrorMsg)
	// UpdateEmailUser(ctx context.Context, recentEmail string, newEmail string) (bool, *models.ErrorMsg)
	// UpdateNameUser(ctx context.Context, user *requests.UserInsertRequest) (bool, *models.ErrorMsg)
	// FindUserByUUID(ctx context.Context, UUID string) (*domain.Users, *models.ErrorMsg)
	// IsEmailRegistered(ctx context.Context, email string) (isRegistered bool, errMsg *models.ErrorMsg)
	// IsNPMRegistered(ctx context.Context, npm string) (isRegistered bool, errMsg *models.ErrorMsg)
}

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

func (U *UsersServiceImplementation) InsertDataUser(ctx context.Context, r *requests.UserInsertRequest) (bool, *response.ErrorMsg) {
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

func (U *UsersServiceImplementation) UpdateDataUser(ctx context.Context, r *requests.UserUpdateRequest) (bool, *response.ErrorMsg) {
	tx, err := U.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}

	response, isEmailRegistered, _ := U.UsersRepository.IsEmailRegistered(ctx, U.DB, r.Email)
	if isEmailRegistered {
		if response.UUID == r.UUID {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_PREVIOUS_FIELD_NOT_ALLOWED, "Email yang sama tidak dapat diubah, gunakan email yang baru.")
		} else {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "Email sudah digunakan.")
		}
	}

	response, isNPMRegistered, _ := U.UsersRepository.IsNPMRegistered(ctx, U.DB, r.NPM)
	if isNPMRegistered {
		if response.UUID == r.UUID {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_PREVIOUS_FIELD_NOT_ALLOWED, "NPM yang sama tidak dapat diubah, gunakan NPM yang baru.")
		} else {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "NPM sudah digunakan.")
		}
	}

	user := &domain.Users{
		UUID:    r.UUID,
		Name:    r.Name,
		Email:   r.Email,
		NPM:     r.NPM,
		ClassID: r.ClassID,
	}

	_, errMsg := U.UsersRepository.UpdateDataUser(ctx, tx, user)
	if errMsg != nil {
		return false, errMsg
	}

	return true, nil
}

// func (U *UsersServiceImplementation) DeleteDataUser(ctx context.Context, UUID string) (bool, *models.ErrorMsg) {
// }

// func (U *UsersServiceImplementation) UpdateEmailUser(ctx context.Context, recentEmail string, newEmail string) (bool, *models.ErrorMsg) {
// }

// func (U *UsersServiceImplementation) UpdateNameUser(ctx context.Context, user *requests.UserInsertRequest) (bool, *models.ErrorMsg) {
// }

// func (U *UsersServiceImplementation) FindUserByUUID(ctx context.Context, UUID string) (*domain.Users, *models.ErrorMsg) {
// }
