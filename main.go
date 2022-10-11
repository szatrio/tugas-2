package main

import (
	"golang_tugas_3/db"
	"golang_tugas_3/server"
	"golang_tugas_3/server/controllers"
	// "github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Orders API
// @description Sample API Spec for Orders
// @version v1.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:4000
// @contact.name Reyhan
// @contact.email reyhan@gmail.com

func main() {
	db := db.ConnectGorm()
	peopleController := controllers.NewPersonController(db)
	router := server.NewRouter(peopleController)
	router.Start(":4000")
}
