package api

import (
	"github.com/dimassfeb-09/sinaustudio.git/repository"
)

type MicroServiceServer interface {
	UserRepository() repository.UsersRepository
	AuthRepository() repository.AuthRepository
	ClassRepository() repository.ClassRepository
	LectureRepository() repository.LectureRepository
	RoomRepository() repository.RoomRepository
	MatkulRepository() repository.MataKuliahRepository
}

type MicroService struct {
	User    repository.UsersRepository
	Auth    repository.AuthRepository
	Class   repository.ClassRepository
	Lecture repository.LectureRepository
	Room    repository.RoomRepository
	Matkul  repository.MataKuliahRepository
}

func NewMicroService(usersRepository repository.UsersRepository, authRepository repository.AuthRepository, classRepository repository.ClassRepository, lectureRepository repository.LectureRepository, roomRepository repository.RoomRepository, matkulRepositoru repository.MataKuliahRepository) MicroServiceServer {
	return &MicroService{User: usersRepository, Auth: authRepository, Class: classRepository, Lecture: lectureRepository, Room: roomRepository, Matkul: matkulRepositoru}
}

func (m *MicroService) UserRepository() repository.UsersRepository {
	return m.User
}

func (m *MicroService) AuthRepository() repository.AuthRepository {
	return m.Auth
}

func (m *MicroService) ClassRepository() repository.ClassRepository {
	return m.Class
}

func (m *MicroService) LectureRepository() repository.LectureRepository {
	return m.Lecture
}

func (m *MicroService) RoomRepository() repository.RoomRepository {
	return m.Room
}

func (m *MicroService) MatkulRepository() repository.MataKuliahRepository {
	return m.Matkul
}
