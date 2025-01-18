package app

import (
	"github.com/hertzcodes/snapp-chat/server/config"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
}
