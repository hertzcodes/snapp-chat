package service

import (
	"github.com/hertzcodes/snapp-chat/server/internal/adapters/storage"
	"github.com/hertzcodes/snapp-chat/server/internal/api/handlers/common"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	service storage.UserRepo
}

func NewUserService(service storage.UserRepo) *UserService {
	return &UserService{service: service}
}

func (u *UserService) SignIn(req common.LoginRequest) (uint, error) {
	user, err := u.service.GetUserByUsername(req.Username)
	if err != nil {
		return 0, err
	}

	isMatched := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if isMatched != nil {
		return 0, err
	}

	return user.ID, nil
}
