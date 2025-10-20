package services

import (
	"users-service/models"
)

type UserService interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}
type Service struct {
	repo UserService
}

func NewUserService(repo UserService) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUserById(id int) (*models.User, error) {
	findId, err := s.repo.GetUserById(id)
	if err != nil {
		return findId, err
	}
	return findId, nil
}

func (s *Service) CreateUser(user *models.User) (*models.User, error) {
	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
