package testcontainer

import (
	"context"
	"fmt"
	"github.com/nopecho/golang-template/test/testdata"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

const (
	postgresImage = "postgres:16-alpine"
	username      = "test"
	password      = "test"
	database      = "test"
)

type PostgresContainer struct {
	Container *postgres.PostgresContainer
	Context   *context.Context
	DSN       string
}

func NewPostgresContainer() *PostgresContainer {
	ctx := context.Background()
	postgresContainer, err := postgres.Run(ctx,
		postgresImage,
		postgres.WithInitScripts(testdata.Postgres),
		postgres.WithDatabase(username),
		postgres.WithUsername(password),
		postgres.WithPassword(database),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		log.Fatal().Msgf("failed to start postgres container: %s", err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		log.Fatal().Msgf("failed to get postgres container host: %s", err)
	}
	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal().Msgf("failed to get postgres container port: %s", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port.Port())

	return &PostgresContainer{
		Container: postgresContainer,
		Context:   &ctx,
		DSN:       dsn,
	}
}

func (p *PostgresContainer) Terminate() {
	if err := p.Container.Terminate(*p.Context); err != nil {
		log.Fatal().Msgf("failed to terminate postgres container: %s", err)
	}
}
