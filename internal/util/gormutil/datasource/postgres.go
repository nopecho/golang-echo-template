package datasource

import (
	"fmt"
	"github.com/nopecho/golang-template/internal/util/common"
	"time"
)

type ConnectionPool struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func defaultConnPool() *ConnectionPool {
	return &ConnectionPool{
		MaxIdleConns:    100,
		MaxOpenConns:    200,
		ConnMaxLifetime: time.Hour,
	}
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	*ConnectionPool
}

func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host:           common.GetEnv("DB_HOST", "localhost"),
		Port:           common.GetEnv("DB_PORT", "5454"),
		Database:       common.GetEnv("DB_DATABASE", "local"),
		Username:       common.GetEnv("DB_USERNAME", "local"),
		Password:       common.GetEnv("DB_PASSWORD", "local"),
		ConnectionPool: defaultConnPool(),
	}
}

func (c *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.Host, c.Username, c.Password, c.Database, c.Port)
}
