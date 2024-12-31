package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"nextgoBlog/model"
)

func Migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		
	}
}

func Tables() []interface{} {
	return []interface{}{
		&model.SessionRecord{},
	}
}