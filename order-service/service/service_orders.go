package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/dto"
	"order-service/model"
)

type OrderServiceInterface interface {
	CreateOrder(order model.Order) (*model.Order, error)
	GetOrderById(id int) (*model.Order, error)
	GetOrderByUserID(userID int) ([]model.Order, error)
}
type OrderService struct {
	repo OrderServiceInterface
}

func NewOrderService(repo OrderServiceInterface) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order model.Order) (*model.Order, error) {
	newOrder, err := s.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (s *OrderService) GetOrderById(id int) (*model.Order, error) {
	newOrder, err := s.repo.GetOrderById(id)
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (s *OrderService) GetOrderByIdUser(req dto.OrderRequest) ([]dto.OrdersResponse, error) {
	// Ambil data user dari users-service
	url := fmt.Sprintf("http://127.0.0.1:8080/users/%d", req.UserID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user not found (status %d)", resp.StatusCode)
	}

	var user dto.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user: %v", err)
	}

	// Ambil order dari DB berdasarkan UserID
	orders, err := s.repo.GetOrderByUserID(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("order not found for user %d: %v", req.UserID, err)
	}

	// Gabungkan hasilnya
	var response []dto.OrdersResponse
	for _, order := range orders {
		response = append(response, dto.OrdersResponse{
			OrderID: order.OrderID,
			Item:    order.ProductName,
			User: dto.UserResponse{
				UserID: user.UserID,
				Name:   user.Name,
				Email:  user.Email,
			},
		})
	}
	return response, nil
}
