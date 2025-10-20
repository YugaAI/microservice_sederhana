package handlers

import (
	"log"
	"strconv"
	"users-service/models"

	"github.com/gin-gonic/gin"
)

type HandlerUser interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}
type UserHandler struct {
	*gin.Engine
	service HandlerUser
}

func NewUserHandler(service HandlerUser) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	newUser, err := h.service.GetUserById(idInt)
	if err != nil {
		c.JSON(400, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, newUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	newUser, err := h.service.CreateUser(&req)
	if err != nil {
		c.JSON(400, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "employee created successfully",
		"data":    newUser,
	})
}
