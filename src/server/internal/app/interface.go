package app

import (
	"github.com/hertzcodes/snapp-chat/server/config"
	"github.com/hertzcodes/snapp-chat/server/internal/api/handlers/service"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	UserService() *service.UserService
}
