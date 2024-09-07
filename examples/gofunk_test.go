package examples

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"github.com/thoas/go-funk"
	"testing"
)

type GoFunkTestSuite struct {
	suite.Suite
	items []TestStruct
}

func (suite *GoFunkTestSuite) TestMap() {
	suite.items = []TestStruct{
		{"a", 1, true},
		{"b", 2, false},
		{"c", 3, true},
	}
	actual := funk.Map(suite.items, func(item TestStruct) string {
		return item.Str
	}).([]string)

	log.Info().Msgf("actual: %v", actual)
	suite.Equal([]string{"a", "b", "c"}, actual)
}

func (suite *GoFunkTestSuite) TestFilter() {
	origin := []TestStruct{
		{"a", 1, true},
		{"b", 2, false},
		{"c", 3, true},
		{"d", 3, false},
		{"e", 3, true},
	}
	suite.items = origin
	actual := funk.Filter(suite.items, func(item TestStruct) bool {
		return !item.Boolean
	}).([]TestStruct)

	suite.Equal([]TestStruct{
		{"b", 2, false},
		{"d", 3, false},
	}, actual)
	suite.Equal(origin, suite.items)
}

func (suite *GoFunkTestSuite) TestFind() {
	suite.items = []TestStruct{
		{"a", 1, true},
		{"b", 2, false},
		{"c", 3, true},
	}
	actual := funk.Find(suite.items, func(item TestStruct) bool {
		return item.Str == "b"
	}).(TestStruct)

	log.Info().Msgf("actual: %v", actual)
	suite.Equal("b", actual.Str)
}

func (suite *GoFunkTestSuite) TestNotFound() {
	suite.items = []TestStruct{
		{"a", 1, true},
		{"b", 2, false},
		{"c", 3, true},
	}
	actual, ok := funk.Find(suite.items, func(item TestStruct) bool {
		return item.Str == "d"
	}).(TestStruct)

	log.Info().Msgf("actual: %v", actual)
	suite.False(ok)
}

func (suite *GoFunkTestSuite) TestEvery() {
	suite.items = []TestStruct{
		{"a", 1, true},
		{"b", 2, true},
		{"c", 3, true},
	}
	actual := funk.Every(suite.items, func(item TestStruct) bool {
		return item.Boolean
	})

	suite.True(actual)
}

func (suite *GoFunkTestSuite) TestSome() {
	suite.items = []TestStruct{
		{"a", 1, true},
		{"b", 2, false},
		{"c", 3, true},
	}
	actual := funk.Some(suite.items, func(item TestStruct) bool {
		return item.Boolean
	})

	suite.True(actual)
}

func (suite *GoFunkTestSuite) Test() {
	suite.items = []TestStruct{
		{"a", 1, true},
		{"b", 2, false},
		{"c", 3, true},
	}

	suite.Equal(3, len(suite.items))
}

func init() {
	setupTestLog()
}

type TestStruct struct {
	Str     string
	Number  int
	Boolean bool
}

func TestGoFunk(t *testing.T) {
	suite.Run(t, new(GoFunkTestSuite))
}
