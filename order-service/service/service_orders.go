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

func (s *OrderService) GetOrderByIdUser(req dto.OrderRequest) (*dto.OrdersResponse, error) {
	url := fmt.Sprintf("http://127.0.0.1:8080/users/%d", req.UserID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user dto.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	//response := &dto.OrdersResponse{
	//	OrderID: req.OrderID,
	//	Item:    req.Item,
	//	User: dto.UserResponse{
	//		Name:   user.Name,
	//		Email:  user.Email,
	//		UserID: user.UserID,
	//	},
	//}
	order := model.Order{
		UserID:      req.UserID,
		ProductName: req.Item,
	}

	if _, err := s.repo.CreateOrder(order); err != nil {
		return nil, err
	}
	return &dto.OrdersResponse{
		OrderID: req.UserID,
		Item:    req.Item,
		User: dto.UserResponse{
			Name:   user.Name,
			Email:  user.Email,
			UserID: user.UserID,
		},
	}, nil
}
