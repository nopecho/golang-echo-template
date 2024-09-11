package database

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64         `json:"id" gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" `
}

type ConnectionInfo struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (c *ConnectionInfo) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.Host, c.Username, c.Password, c.Database, c.Port)
}
