package handler

import "github.com/gin-gonic/gin"

type Dependensi struct {
	OrderHandler *OrderHandler
}

func RegistRoute(r *gin.Engine, deps *Dependensi) {
	ord := r.Group("/orders")
	{
		ord.POST("/create", deps.OrderHandler.CreateOrder)
		ord.GET("/:order_id", deps.OrderHandler.GetOrderbyId)
		ord.GET("/user/:user_id", deps.OrderHandler.GetDetailOrderbyIdUser)

	}
}
