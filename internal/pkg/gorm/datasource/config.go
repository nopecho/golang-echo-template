package datasource

import (
	"fmt"
	"github.com/nopecho/golang-template/internal/pkg/apputil"
	"time"
)

type ConnectionPool struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func DefaultConnPool() *ConnectionPool {
	return &ConnectionPool{
		MaxIdleConns:    10,
		MaxOpenConns:    20,
		ConnMaxLifetime: time.Hour,
	}
}

type ConnectionInfo struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	*ConnectionPool
}

func DefaultConnectionInfo() *ConnectionInfo {
	return &ConnectionInfo{
		Host:           apputil.GetDefaultEnv("DB_HOST", "localhost"),
		Port:           apputil.GetDefaultEnv("DB_PORT", "5454"),
		Database:       apputil.GetDefaultEnv("DB_DATABASE", "local"),
		Username:       apputil.GetDefaultEnv("DB_USERNAME", "local"),
		Password:       apputil.GetDefaultEnv("DB_PASSWORD", "local"),
		ConnectionPool: DefaultConnPool(),
	}
}

// DSN returns the data source name for the connection
func (c *ConnectionInfo) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.Host, c.Username, c.Password, c.Database, c.Port)
}
