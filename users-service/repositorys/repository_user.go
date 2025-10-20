package repositorys

import (
	"log"
	"users-service/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *Repository {
	if db == nil {
		log.Println("⚠️  DB instance is nil in NewRepositoryEmployee")
	} else {
		sqlDB, _ := db.DB()
		if err := sqlDB.Ping(); err != nil {
			log.Printf("DB not connected: %v", err)
		} else {
			log.Println("DB connected successfully")
		}
	}
	return &Repository{db: db}
}

func (r *Repository) GetUserById(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
