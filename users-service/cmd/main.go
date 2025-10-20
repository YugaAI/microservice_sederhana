package main

import (
	"users-service/handlers"
	"users-service/insfrastruktur/configs"
	"users-service/insfrastruktur/setups"
	"users-service/models"
	"users-service/repositorys"
	"users-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()
	db := setups.ConnectDB(cfg)
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	userRepo := repositorys.NewRepositoryUser(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	handlers.RegistrasiRouter(r, &handlers.Dependencies{
		UsrHandler: userHandler,
	})

	r.Run(":8080")
}
