package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) (init bool, err error) {
	m := gormigrate.New(db, &gormigrate.Options{
		TableName: "migrations",
		IDColumnName: "id",
		IDColumnSize: 255,
		UseTransaction: false,
		ValidateUnknownMigrations: true,
	}, Migrations())

	m.InitSchema(func(tx *gorm.DB) error {
		init = true
		return tx.AutoMigrate(Tables()...)
	})

	err = m.Migrate()
	return
}