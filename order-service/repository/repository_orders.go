package repository

import (
	"order-service/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(order model.Order) (*model.Order, error) {
	if err := r.DB.Create(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetOrderById(id int) (*model.Order, error) {
	var order model.Order
	if err := r.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *OrderRepository) GetOrderByUserID(userID int) ([]model.Order, error) {
	var order []model.Order
	if err := r.DB.Where("user_id = ?", userID).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
