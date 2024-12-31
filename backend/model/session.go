package model

import (
	"time"
)

type SessionRecord struct {
	Token     string    `gorm:"type:varchar(255);primaryKey"`
	UserID    string    `gorm:"type:varchar(255);unique"`
	CreatedAt time.Time `gorm:"precision:6"`
}
