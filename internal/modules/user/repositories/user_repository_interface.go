package repositories

import (
	userModel "WriteWise/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
}
