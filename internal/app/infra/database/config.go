package database

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64         `json:"id" gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
