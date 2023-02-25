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
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthService interface {
	AuthRegisterUser(ctx context.Context, r *requests.AuthRegisterRequest) (bool, *response.ErrorMsg)
	AuthLoginUser(ctx context.Context, email string, password string) (*response.UserInfoLogin, *response.ErrorMsg)
}

type AuthRepositoryImplementation struct {
	DB                *sql.DB
	AuthRepository    repository.AuthRepository
	ClassRepository   repository.ClassRepository
	LectureRepository repository.LectureRepository
	M                 api.MicroServiceServer
}

func NewAuthServiceImplementation(DB *sql.DB, M api.MicroServiceServer) AuthService {
	return &AuthRepositoryImplementation{
		DB:                DB,
		AuthRepository:    M.AuthRepository(),
		ClassRepository:   M.ClassRepository(),
		LectureRepository: M.LectureRepository(),
		M:                 M,
	}
}

func (a *AuthRepositoryImplementation) AuthRegisterUser(ctx context.Context, r *requests.AuthRegisterRequest) (bool, *response.ErrorMsg) {
	tx, err := a.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	hashPassword, err := helpers.HashAndSaltPassword([]byte(r.Password))
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}

	_, isClassIDValid, _ := a.M.ClassRepository().FindClassByID(ctx, a.DB, r.ClassID)
	if !isClassIDValid {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Kode kelas tidak ditemukan")
	}

	_, isEmailRegistered, _ := a.M.UserRepository().IsEmailRegistered(ctx, a.DB, r.Email)
	if isEmailRegistered {
		return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "Email telah digunakan")
	}

	user := &domain.AuthRegisterUser{
		Name:     r.Name,
		Email:    r.Email,
		Password: hashPassword,
		Role:     r.Role,
		ClassID:  r.ClassID,
	}

	isRegisterSuccess, lastID, errMsg := a.AuthRepository.AuthRegisterUser(ctx, tx, user)
	if errMsg != nil && !isRegisterSuccess {
		return false, errMsg
	}

	if r.Role == "dosen" || r.Role == "guru" {
		if isRegisterSuccess {
			lecture := &domain.Lecture{
				Name:   user.Name,
				UserID: lastID,
			}
			_, errMsg := a.LectureRepository.InsertLecture(ctx, tx, lecture)
			if errMsg != nil {
				return false, helpers.ToErrorMsg(http.StatusBadRequest, exception.ERR_BAD_REQUEST_FIELD, "User ID telah digunakan")
			}
		}
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
		Role:    userResponse.Role,
		ClassID: userResponse.ClassID,
	}

	return userInfo, nil
}
