package app

import (
	"errors"

	"github.com/hertzcodes/snapp-chat/server/config"
	"github.com/hertzcodes/snapp-chat/server/internal/adapters/postgres"
	"github.com/hertzcodes/snapp-chat/server/internal/adapters/storage/entities"
	"gorm.io/gorm"
)

type app struct {
	db  *gorm.DB
	cfg config.Config
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entities.User{},
		&entities.Room{},

		// ... other tables ...
	); err != nil {
		return errors.New("migration failed")
	}
	return nil
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	if err := migrate(db); err != nil {
		return err
	}
	a.db = db
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)

	if err != nil {
		panic(err)
	}
	return app
}
