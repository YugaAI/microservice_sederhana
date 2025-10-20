package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"order-service/dto"
	"order-service/model"
	"strconv"
)

type HandlerOrder interface {
	CreateOrder(order model.Order) (*model.Order, error)
	GetOrderById(id int) (*model.Order, error)
	GetOrderByIdUser(req dto.OrderRequest) ([]dto.OrdersResponse, error)
}
type OrderHandler struct {
	*gin.Engine
	service HandlerOrder
}

func NewOrderHandler(service HandlerOrder) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah user ada di users-service
	userURL := fmt.Sprintf("http://localhost:8080/users/%d", req.UserID)
	response, err := http.Get(userURL)
	if err != nil || response.StatusCode != 200 {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}
	defer response.Body.Close()

	var user map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		c.JSON(500, gin.H{"error": "failed to decode user response"})
		return
	}

	// Simpan order ke database melalui service
	createdOrder, err := h.service.CreateOrder(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan hasil
	c.JSON(200, gin.H{
		"status": "success",
		"order":  createdOrder,
		"user":   user,
	})
}

func (h *OrderHandler) GetOrderbyId(c *gin.Context) {
	orderId := c.Param("order_id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	order, _ := h.service.GetOrderById(orderIdInt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"data": order, "status": "success"})
}

func (h *OrderHandler) GetDetailOrderbyIdUser(c *gin.Context) {
	userId := c.Param("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := dto.OrderRequest{
		UserID: userIdInt,
	}

	order, err := h.service.GetOrderByIdUser(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"data": order, "status": "success"})
}
