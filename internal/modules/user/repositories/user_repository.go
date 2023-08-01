package repositories

import (
	userModel "WriteWise/internal/modules/user/models"
	"WriteWise/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (ur *UserRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User

	ur.DB.Create(&user).Scan(&newUser)

	return newUser
}
