package gorm_migrations

import (
	"github.com/labstack/gommon/log"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/entities"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/migrations"
	"gorm.io/gorm"
)

type gormMigrations struct {
	db       *gorm.DB
	entities entities.Entities
}

func NewMigrator(db *gorm.DB, entities entities.Entities) migrations.Migrations {
	return &gormMigrations{db: db, entities: entities}
}

func (g *gormMigrations) Run() {
	if err := g.db.AutoMigrate(g.entities...); err != nil {
		log.Errorf("Failed during auto migrations with gorm, %s", err.Error())
	}
}
