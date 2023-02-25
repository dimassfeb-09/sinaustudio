package main

import (
	"database/sql"
	"github.com/dimassfeb-09/sinaustudio.git/api"
	"github.com/dimassfeb-09/sinaustudio.git/controllers"
	"github.com/dimassfeb-09/sinaustudio.git/repository"
	"github.com/dimassfeb-09/sinaustudio.git/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db := api.ConnectionDatabases()
	defer db.Close()

	route := gin.Default()
	route.Use(api.ControllAccessAllow())

	Router(route, db)
}

func Router(route *gin.Engine, db *sql.DB) {
	v1 := route.Group("/api/v.1/")

	usersRepository := repository.NewUsersRepositoryImplementations()
	authRepository := repository.NewAuthRepositoryImplementation()
	classRepository := repository.NewClassRepositoryImplementation()
	lectureRepository := repository.NewLectureRepositoryImplementation()
	roomRepository := repository.NewRoomRepositoryImplementation()
	matkulRepository := repository.NewMataKuliahRepositoryImplementation()

	microServices := api.NewMicroService(usersRepository, authRepository, classRepository, lectureRepository, roomRepository, matkulRepository)

	usersService := services.NewUserServiceImplementation(db, microServices)
	authService := services.NewAuthServiceImplementation(db, microServices)
	classService := services.NewClassServiceImplementation(db, microServices)
	lectureService := services.NewLectureServiceImplementation(db, microServices)
	roomService := services.NewRoomServiceImplementation(db, microServices)
	matkulService := services.NewMataKuliahServiceImplementation(db, microServices)

	usersController := controllers.NewUsersControllerImplementation(usersService)
	authController := controllers.NewAuthControllerImplementation(authService)
	classController := controllers.NewClassControllerImplementation(classService)
	lectureController := controllers.NewLectureController(lectureService)
	roomController := controllers.NewRoomController(roomService)
	matkulController := controllers.NewMatkulController(matkulService)

	auth := v1.Group("/auth")
	auth.POST("/register", authController.AuthRegisterUser)
	auth.POST("/login", authController.AuthLoginUser)

	v1.Use(api.MiddlewareAuthorization)
	user := v1.Group("/user")
	user.POST("/create", usersController.InsertDataUser)
	user.PUT("/update", usersController.UpdateDataUser)
	user.DELETE("/delete", usersController.DeleteDataUser)
	user.PUT("/changepassword", usersController.ChangePasswordUser)

	class := v1.Group("/class")
	class.POST("/create", classController.AddClass)
	class.PUT("/update", classController.UpdateClass)
	class.DELETE("/delete", classController.DeleteClassByID)
	class.GET("/", func(c *gin.Context) {
		if ID := c.Query("id"); ID != "" {
			classController.FindClassByID(c)
			return
		} else if name := c.Query("name"); name != "" {
			classController.FindClassByName(c)
			return
		}
	})

	lecture := v1.Group("/lecture")
	lecture.POST("/create", lectureController.InsertLecture)
	lecture.PUT("/update", lectureController.UpdateLecture)
	lecture.DELETE("/delete", lectureController.DeleteLecture)
	lecture.GET("/", func(c *gin.Context) {
		if id := c.Query("id"); id != "" {
			lectureController.FindLectureByID(c)
			return
		} else if name := c.Query("name"); name != "" {
			lectureController.FindLectureByName(c)
			return
		}
	})

	room := v1.Group("/room")
	room.POST("/create", roomController.InsertRoom)
	room.PUT("/update", roomController.UpdateRoom)
	room.DELETE("/delete", roomController.DeleteRoom)
	room.GET("/", func(c *gin.Context) {
		if id := c.Query("id"); id != "" {
			roomController.FindRoomByID(c)
			return
		}
	})

	matkul := v1.Group("/matkul")
	matkul.POST("/create", matkulController.InsertMatkul)
	matkul.PUT("/update", matkulController.UpdateMatkul)
	matkul.DELETE("/delete", matkulController.DeleteMatkul)
	matkul.GET("/", func(c *gin.Context) {
		if id := c.Query("id"); id != "" {
			matkulController.FindMatkulByID(c)
			return
		} else if name := c.Query("name"); name != "" {
			matkulController.FindMatkulByName(c)
			return
		}
	})

	err := route.Run(":8081")
	if err != nil {
		log.Fatalln(err)
	}
}
