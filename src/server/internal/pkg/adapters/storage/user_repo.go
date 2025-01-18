package storage

import (
	"errors"

	"github.com/hertzcodes/snapp-chat/server/internal/pkg/adapters/storage/entities"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *entities.User) error {
	if user == nil {
		return errors.New("nil user")
	}
	return r.db.Table("users").Create(entities.User{}).Error
}

func (r *userRepo) Update(user *entities.User) error {
	if user == nil {
		return errors.New("nil user")
	}

	return r.db.Table("users").Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepo) Delete(id uint) error {
	return r.db.Table("users").Delete(&entities.User{}, id).Error
}

func (r *userRepo) GetUserByID(id uint) (*entities.User, error) {
	var user entities.User

	q := r.db.Table("users").Where("id = ?", id)

	err := q.First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// func (r *userRepo) GetUserByUsername(username string) (*entities.User, error) {
// 	var user entities.User

// 	q := r.db.Table("users").Where("username = ?", username)

// 	err := q.First(&user).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
