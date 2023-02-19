package main

import (
	"github.com/dimassfeb-09/sinaustudio.git/app"
	"github.com/dimassfeb-09/sinaustudio.git/controllers"
	"github.com/dimassfeb-09/sinaustudio.git/middleware"
	"github.com/dimassfeb-09/sinaustudio.git/repository"
	"github.com/dimassfeb-09/sinaustudio.git/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db := app.ConnectionDatabases()
	defer db.Close()

	route := gin.Default()
	route.Use(middleware.ControllAccessAllow())

	v1 := route.Group("/api/v.1/")

	usersRepository := repository.NewUsersRepositoryImplementations()
	authRepository := repository.NewAuthRepositoryImplementation()
	classRepository := repository.NewClassRepositoryImplementation()
	lectureRepository := repository.NewLectureRepositoryImplementation()

	microServices := app.NewMicroService(usersRepository, authRepository, classRepository, lectureRepository)

	usersService := services.NewUserServiceImplementation(db, microServices)
	authService := services.NewAuthServiceImplementation(db, microServices)
	classService := services.NewClassServiceImplementation(db, microServices)
	lectureService := services.NewLectureServiceImplementation(db, microServices)

	usersController := controllers.NewUsersControllerImplementation(usersService)
	authController := controllers.NewAuthControllerImplementation(authService)
	classController := controllers.NewClassControllerImplementation(classService)
	lectureController := controllers.NewLectureController(lectureService)

	auth := v1.Group("/auth")
	auth.POST("/register", authController.AuthRegisterUser)
	auth.POST("/login", authController.AuthLoginUser)

	v1.Use(middleware.MiddlewareAuthorizationfunc)
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
	lecture.DELETE("/", lectureController.DeleteLecture)
	lecture.GET("/", func(c *gin.Context) {
		if id := c.Query("id"); id != "" {
			lectureController.FindLectureByID(c)
			return
		} else if name := c.Query("name"); name != "" {
			lectureController.FindLectureByName(c)
			return
		}
	})

	err := route.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
