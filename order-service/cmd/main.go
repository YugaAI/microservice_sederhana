package main

import (
	"order-service/handler"
	"order-service/infrastructure/bootstrap"
	"order-service/infrastructure/config"
	"order-service/model"
	"order-service/repository"
	"order-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := bootstrap.ConnectDB(cfg)
	db.AutoMigrate(&model.Order{})

	r := gin.Default()

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	handler.RegistRoute(r, &handler.Dependensi{
		OrderHandler: orderHandler,
	})

	r.Run(":8085")
}
