package main

import (
	"github.com/dimassfeb-09/sinaustudio.git/app"
	"github.com/dimassfeb-09/sinaustudio.git/controllers"
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
	v1 := route.Group("/api/v.1/")

	usersRepository := repository.NewUsersRepositoryImplementations()
	usersService := services.NewUserServiceImplementation(db, usersRepository)
	usersController := controllers.NewUsersControllerImplementation(usersService)

	v1.POST("/user", usersController.InsertDataUser)
	v1.PUT("/user/", usersController.UpdateDataUser)
	v1.DELETE("/user/", usersController.DeleteDataUser)

	err := route.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}

}
