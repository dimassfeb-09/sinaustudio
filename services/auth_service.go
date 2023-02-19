package services

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/sinaustudio.git/app"
	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthService interface {
	AuthRegisterUser(ctx context.Context, r *requests.AuthRegisterRequest) (bool, *response.ErrorMsg)
	AuthLoginUser(ctx context.Context, email string, password string) (*response.UserInfoLogin, *response.ErrorMsg)
}

type AuthRepositoryImplementation struct {
	DB             *sql.DB
	AuthRepository repository.AuthRepository
	M              app.MicroServiceServer
}

func NewAuthServiceImplementation(DB *sql.DB, M app.MicroServiceServer) AuthService {
	return &AuthRepositoryImplementation{DB: DB, AuthRepository: M.AuthRepository(), M: M}
}

func (a *AuthRepositoryImplementation) AuthRegisterUser(ctx context.Context, r *requests.AuthRegisterRequest) (bool, *response.ErrorMsg) {
	tx, err := a.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	password, err := helpers.HashAndSaltPassword([]byte(r.Password))
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}

	userRepository := a.M.UserRepository()

	_, isEmailRegistered, _ := userRepository.IsEmailRegistered(ctx, a.DB, r.Email)
	if isEmailRegistered {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Email telah digunakan")
	}

	if r.NPM != "" {
		_, isNPMRegistered, _ := userRepository.IsNPMRegistered(ctx, a.DB, r.NPM)
		if isNPMRegistered {
			return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "NPM telah digunakan")
		}
	}

	r.Password = password
	user := &domain.AuthRegisterUser{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Role:     r.Role,
		NPM:      r.NPM,
		ClassID:  r.ClassID,
	}

	success, errMsg := a.AuthRepository.AuthRegisterUser(ctx, tx, user)
	if errMsg != nil && !success {
		return false, errMsg
	}

	return true, nil
}

func (a *AuthRepositoryImplementation) AuthLoginUser(ctx context.Context, email string, password string) (*response.UserInfoLogin, *response.ErrorMsg) {

	_, isEmailRegistered, _ := a.M.UserRepository().IsEmailRegistered(ctx, a.DB, email)
	if !isEmailRegistered {
		return nil, helpers.ToErrorMsg(http.StatusNotFound, exception.ERR_NOT_FOUND, "Email tidak ditemukan.")
	}

	success, result, errMsg := a.AuthRepository.AuthLoginUser(ctx, a.DB, email)
	if errMsg != nil && !success {
		return nil, errMsg
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return nil, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Password anda salah. Silahkan periksa kembali.")
	}

	userResponse, errMsg := a.M.UserRepository().FindUserByEmail(ctx, a.DB, email)
	if err != nil {
		return nil, errMsg
	}

	userInfo := &response.UserInfoLogin{
		ID:      userResponse.ID,
		Name:    userResponse.Name,
		Email:   userResponse.Email,
		NPM:     userResponse.NPM,
		Role:    userResponse.Role,
		ClassID: userResponse.ClassID,
	}

	return userInfo, nil
}
