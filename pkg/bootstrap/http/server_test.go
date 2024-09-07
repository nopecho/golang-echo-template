package http

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("MockHandler.ServeHTTP: %v", r)
	m.Called(w, r)
}

type HttpServerTestSuite struct {
	suite.Suite
	handler *MockHandler
}

func (suite *HttpServerTestSuite) TestNewHttpServer() {
	mux := http.NewServeMux()

	mux.Handle("xx", nil)
	//actual := NewHttpServer(8080, nil)
	//
	//suite.Equal(8080, actual.port)
	//suite.Nil(actual.handler)
	//suite.NotNil(actual.server)
}

func (suite *HttpServerTestSuite) TestStart() {
	//server := NewHttpServer(8080, suite.handler)
	//go server.Start()
	//time.Sleep(1 * time.Second)
	//
	//suite.handler.AssertCalled(suite.T(), "ServeHTTP", nil, nil)
}

func TestHttpServer(t *testing.T) {
	suite.Run(t, new(HttpServerTestSuite))
}
