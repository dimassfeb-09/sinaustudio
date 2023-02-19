package services

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/sinaustudio.git/app"
	"golang.org/x/crypto/bcrypt"
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
	M               app.MicroServiceServer
}

func NewUserServiceImplementation(DB *sql.DB, M app.MicroServiceServer) UsersService {
	return &UsersServiceImplementation{
		DB:              DB,
		UsersRepository: M.UserRepository(),
		M:               M,
	}
}

type UsersService interface {
	InsertDataUser(ctx context.Context, r *requests.UserInsertRequest) (bool, *responseError.ErrorMsg)
	UpdateDataUser(ctx context.Context, r *requests.UserUpdateRequest) (bool, *responseError.ErrorMsg)
	DeleteDataUser(ctx context.Context, confirmPass string, ID int) (bool, *responseError.ErrorMsg)
	FindUserByID(ctx context.Context, ID int) (*domain.Users, *responseError.ErrorMsg)
	IsEmailRegistered(ctx context.Context, email string) (isRegistered bool, errMsg *responseError.ErrorMsg)
	IsNPMRegistered(ctx context.Context, npm string) (isRegistered bool, errMsg *responseError.ErrorMsg)
	ChangePasswordUser(ctx context.Context, ID int, recentPass string, newPass string) (isSuccess bool, errNsg *responseError.ErrorMsg)
}

func (U *UsersServiceImplementation) InsertDataUser(ctx context.Context, r *requests.UserInsertRequest) (bool, *responseError.ErrorMsg) {
	tx, err := U.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err)
	}
	defer helpers.RollbackOrCommit(tx)

	password, err := helpers.HashAndSaltPassword([]byte(r.Password))
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}

	r.Password = password
	_, isIDRegistered, _ := U.UsersRepository.FindUserByID(ctx, U.DB, r.ID)
	if isIDRegistered {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "ID sudah terdaftar")
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
		ID:       r.ID,
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Role:     r.Role,
		NPM:      r.NPM,
		ClassID:  r.ClassID,
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
		if response.ID == r.ID {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_PREVIOUS_FIELD_NOT_ALLOWED, "Gunakan email yang baru.")
		} else {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "Email sudah digunakan.")
		}
	}

	response, isNPMRegistered, _ := U.UsersRepository.IsNPMRegistered(ctx, U.DB, r.NPM)
	if r.NPM != "" {
		if isNPMRegistered {
			if response.ID == r.ID {
				return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_PREVIOUS_FIELD_NOT_ALLOWED, "Gunakan NPM yang baru.")
			} else {
				return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_ALREADY_USE, "NPM sudah digunakan.")
			}
		}
	}

	user := &domain.Users{
		ID:      r.ID,
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

func (U *UsersServiceImplementation) DeleteDataUser(ctx context.Context, confirmPass string, ID int) (bool, *responseError.ErrorMsg) {
	tx, err := U.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	user, isUserRegistered, errMsg := U.UsersRepository.FindUserByID(ctx, U.DB, ID)
	if !isUserRegistered {
		return false, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "ID tidak ditemukan.")
	}

	if isUserRegistered {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(confirmPass))
		if err != nil {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Password tidak sesuai.")
		}

		isSuccess, errMsg := U.UsersRepository.DeleteDataUser(ctx, tx, user.ID)
		if errMsg != nil && !isSuccess {
			return false, errMsg
		}

		return isSuccess, nil
	} else {
		return false, errMsg
	}
}

func (U *UsersServiceImplementation) FindUserByID(ctx context.Context, ID int) (*domain.Users, *responseError.ErrorMsg) {
	response, isIDRegistered, errMsg := U.UsersRepository.FindUserByID(ctx, U.DB, ID)
	if errMsg != nil {
		return nil, errMsg
	}

	if !isIDRegistered {
		return nil, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Data ID tidak ditemukan.")
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

func (U *UsersServiceImplementation) ChangePasswordUser(ctx context.Context, ID int, recentPass string, newPass string) (isSuccess bool, errNsg *responseError.ErrorMsg) {
	tx, err := U.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	user, isUserRegistered, errMsg := U.UsersRepository.FindUserByID(ctx, U.DB, ID)
	if !isUserRegistered && errMsg != nil {
		return false, errMsg
	}

	if isUserRegistered {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(recentPass))
		if err != nil {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Password saat ini tidak sesuai.")
		}

		hashNewPass, err := helpers.HashAndSaltPassword([]byte(newPass))
		if err != nil {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, err)
		}

		isSuccess, errMsg := U.UsersRepository.ChangePasswordUser(ctx, tx, hashNewPass, user.ID)
		if errMsg != nil && !isSuccess {
			return false, errMsg
		}

		return isSuccess, nil
	} else {
		return false, errMsg
	}
}
