package synk

import (
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type ChannelManagerTestSuite struct {
	suite.Suite
}

func (suite *ChannelManagerTestSuite) TestOpenBufferChannel() {
	actual := OpenBufferChannel[string](10)
	suite.NotNil(actual)
	suite.NotNil(actual.Channel)
	suite.NotNil(actual.wg)
}

func (suite *ChannelManagerTestSuite) TestSend() {
	size := 10
	cm := OpenBufferChannel[string](size)

	for i := range size {
		go cm.Send(strconv.Itoa(i))
	}
	actual := cm.ReceiveWait()

	suite.Len(actual, size)
}

func TestChannelManager(t *testing.T) {
	suite.Run(t, new(ChannelManagerTestSuite))
}
