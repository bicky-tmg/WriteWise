package services

import (
	"WriteWise/internal/modules/user/requests/auth"
	UserResponse "WriteWise/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
}
