package examples

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestService interface {
	DoSomething(input interface{}) (interface{}, error)
}

type MockService struct {
	mock.Mock
}

func (m *MockService) DoSomething(input interface{}) (interface{}, error) {
	args := m.Called(input)
	return args.Get(0), args.Error(1)
}

type MockServiceTestSuite struct {
	suite.Suite
	service TestService
}

func (suite *MockServiceTestSuite) SetupTest() {
	suite.service = new(MockService)
	log.Info().Msg("setup mock service")
}

func (suite *MockServiceTestSuite) TestMock() {
	sut := suite.service.(*MockService)
	sut.On("DoSomething", "input").Return("output", nil)

	actual, err := suite.service.DoSomething("input")

	sut.AssertNumberOfCalls(suite.T(), "DoSomething", 1)
	sut.AssertCalled(suite.T(), "DoSomething", "input")
	suite.Equal("output", actual)
	suite.NoError(err)
}

func (suite *MockServiceTestSuite) TestMockNotCalled() {
	sut := suite.service.(*MockService)
	sut.On("DoSomething", "input").Return("output", nil)

	sut.AssertNotCalled(suite.T(), "DoSomething")
	sut.AssertNumberOfCalls(suite.T(), "DoSomething", 0)
}

func (suite *MockServiceTestSuite) TestMockType() {
	service := suite.service

	suite.IsType(service, &MockService{})
}

func TestTestify(t *testing.T) {
	suite.Run(t, new(MockServiceTestSuite))
}
