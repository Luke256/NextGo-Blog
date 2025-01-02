package model

import (
	"bytes"
	"time"
	"encoding/gob"
)

type SessionRecord struct {
	Token     string    `gorm:"type:varchar(255);primaryKey"`
	UserID    string    `gorm:"type:varchar(255);unique"`
	CreatedAt time.Time `gorm:"precision:6"`
	Data      []byte    `gorm:"type:longblob"`
}

func (SessionRecord) TableName() string {
	return "r_sessions"
}

func (s *SessionRecord) GetData() (data map[string]interface{}, err error) {
	return data, gob.NewDecoder(bytes.NewBuffer(s.Data)).Decode(&data)
}

func (s *SessionRecord) SetData(data map[string]interface{}) {
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(data)
	if err != nil {
		panic(err)
	}
	s.Data = buf.Bytes()
}