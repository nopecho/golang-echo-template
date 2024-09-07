package config

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type EnvConfigTestSuite struct {
	suite.Suite
}

func (suite *EnvConfigTestSuite) TestGetOrDefaultEnv() {
	actual := getOrDefaultEnv("TEST_PORT", "10000")

	suite.Equal("10000", actual)
}

func (suite *EnvConfigTestSuite) TestGetEnv() {
	_ = os.Setenv("TEST_PORT", "20000")
	defer os.Unsetenv("TEST_PORT")

	actual := getOrDefaultEnv("TEST_PORT", "10000")

	suite.Equal("20000", actual)
}

func TestEnvConfig(t *testing.T) {
	suite.Run(t, new(EnvConfigTestSuite))
}
