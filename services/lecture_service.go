package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimassfeb-09/sinaustudio.git/app"
	"github.com/dimassfeb-09/sinaustudio.git/entity/domain"
	"github.com/dimassfeb-09/sinaustudio.git/entity/requests"
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/dimassfeb-09/sinaustudio.git/exception"
	"github.com/dimassfeb-09/sinaustudio.git/helpers"
	"github.com/dimassfeb-09/sinaustudio.git/repository"
	"net/http"
)

type LectureService interface {
	InsertLecture(ctx context.Context, r *requests.InsertLectureRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	UpdateLecture(ctx context.Context, r *requests.UpdateLectureRequest) (isSuccess bool, errMsg *response.ErrorMsg)
	DeleteLectureByID(ctx context.Context, ID int) (isSuccess bool, errMsg *response.ErrorMsg)
	FindLectureByID(ctx context.Context, ID int) (r *response.LectureResponse, isValid bool, errMsg *response.ErrorMsg)
	FindLectureByName(ctx context.Context, name string) (r *response.LectureResponse, isValid bool, errMsg *response.ErrorMsg)
}

type LectureServiceImplementation struct {
	DB                *sql.DB
	LectureRepository repository.LectureRepository
	M                 app.MicroServiceServer
}

func NewLectureServiceImplementation(DB *sql.DB, m app.MicroServiceServer) LectureService {
	return &LectureServiceImplementation{DB: DB, LectureRepository: m.LectureRepository(), M: m}
}

func (l *LectureServiceImplementation) InsertLecture(ctx context.Context, r *requests.InsertLectureRequest) (bool, *response.ErrorMsg) {
	tx, err := l.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isIDValid, errMsg := l.LectureRepository.FindLectureByID(ctx, l.DB, r.ID)
	if errMsg != nil && !isIDValid {
		return false, errMsg
	}

	isSuccess, errMsg := l.LectureRepository.InsertLecture(ctx, tx, r.Name)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return true, nil
}

func (l *LectureServiceImplementation) UpdateLecture(ctx context.Context, r *requests.UpdateLectureRequest) (bool, *response.ErrorMsg) {
	tx, err := l.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	_, isIDValid, errMsg := l.LectureRepository.FindLectureByID(ctx, l.DB, r.ID)
	if errMsg != nil && !isIDValid {
		return false, errMsg
	}

	lecture := &domain.Lecture{
		ID:   r.ID,
		Name: r.Name,
	}

	isSuccess, errMsg := l.LectureRepository.UpdateLecture(ctx, tx, lecture)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return true, nil
}

func (l *LectureServiceImplementation) DeleteLectureByID(ctx context.Context, ID int) (bool, *response.ErrorMsg) {
	tx, err := l.DB.Begin()
	if err != nil {
		return false, helpers.ToErrorMsg(http.StatusInternalServerError, exception.ERR_INTERNAL_SERVER, err)
	}
	defer helpers.RollbackOrCommit(tx)

	isSuccess, errMsg := l.LectureRepository.DeleteLectureByID(ctx, tx, ID)
	if errMsg != nil && !isSuccess {
		return false, errMsg
	}

	return true, nil
}

func (l *LectureServiceImplementation) FindLectureByID(ctx context.Context, ID int) (r *response.LectureResponse, isValid bool, errMsg *response.ErrorMsg) {
	lecture, isIDValid, errMsg := l.LectureRepository.FindLectureByID(ctx, l.DB, ID)
	if isIDValid {
		fmt.Println(lecture)
		lectureResponse := &response.LectureResponse{
			ID:   lecture.ID,
			Name: lecture.Name,
		}
		return lectureResponse, true, nil
	} else {
		return nil, false, errMsg
	}
}

func (l *LectureServiceImplementation) FindLectureByName(ctx context.Context, name string) (r *response.LectureResponse, isValid bool, errMsg *response.ErrorMsg) {
	lecture, isIDValid, errMsg := l.LectureRepository.FindLectureByName(ctx, l.DB, name)
	if isIDValid {
		lectureResponse := &response.LectureResponse{
			ID:   lecture.ID,
			Name: lecture.Name,
		}
		return lectureResponse, true, nil
	} else {
		return nil, false, errMsg
	}
}
