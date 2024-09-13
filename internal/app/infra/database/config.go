package database

import (
	"fmt"
	"github.com/nopecho/golang-template/internal/pkg/helper"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64         `json:"id" gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ConnectionInfo struct {
	host     string
	port     string
	database string
	username string
	password string
}

func EnvConnectionInfo() *ConnectionInfo {
	return &ConnectionInfo{
		host:     helper.GetDefaultEnv("POSTGRES_HOST", "localhost"),
		port:     helper.GetDefaultEnv("POSTGRES_PORT", "5454"),
		database: helper.GetDefaultEnv("POSTGRES_DATABASE", "local"),
		username: helper.GetDefaultEnv("POSTGRES_USERNAME", "local"),
		password: helper.GetDefaultEnv("POSTGRES_PASSWORD", "local"),
	}
}

// DSN returns the data source name for the connection
func (c *ConnectionInfo) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.host, c.username, c.password, c.database, c.port)
}
