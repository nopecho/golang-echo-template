package datasource

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	stdlog "log"
	"time"
)

func NewPostgres(config *PostgresConfig) *gorm.DB {
	return NewPostgresWith(config.DSN(), config.ConnectionPool)
}

func NewDefaultPostgres() *gorm.DB {
	return NewPostgres(NewPostgresConfig())
}

func NewPostgresWith(dsn string, pool *ConnectionPool) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		Logger: logger.New(
			stdlog.Default(),
			logger.Config{
				SlowThreshold:             2 * time.Second,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		log.Fatal().Msgf("postgresql is not connected. %v", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal().Msgf("postgresql is not connected. %v", err)
	}

	if pool == nil {
		pool = defaultConnPool()
	}

	if pool.ConnMaxLifetime <= 0 {
		pool.ConnMaxLifetime = time.Hour
	}

	postgresDB.SetMaxIdleConns(pool.MaxIdleConns)
	postgresDB.SetMaxOpenConns(pool.MaxOpenConns)
	postgresDB.SetConnMaxLifetime(pool.ConnMaxLifetime)
	return db
}

func AutoMigrate(db *gorm.DB, models ...interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatal().Msgf("auto migration failed. %v", err)
	}
}
