package gorm

import (
	"nextgoBlog/migration"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB, doMigration bool) (repo *Repository, init bool, err error) {
	repo = &Repository{
		DB: db,
	}

	if doMigration {
		init, err = migration.Migrate(db)
		if err != nil {
			return nil, false, err
		}
	}

	return 
}