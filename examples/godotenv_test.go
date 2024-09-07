package examples

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type DotEnvTestSuite struct {
	suite.Suite
}

func (suite *DotEnvTestSuite) TestLoad() {
	suite.Run("Load Env Test", func() {
		/**
		 * ./.env.example
		 * TEST_ENV=foo
		 */
		_ = godotenv.Load("./.env.example")

		actual := os.Getenv("TEST_ENV")

		suite.Equal("foo", actual)
	})
}

func TestDotEnv(t *testing.T) {
	suite.Run(t, new(DotEnvTestSuite))
}
