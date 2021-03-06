package message_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ysugimoto/gqtt/message"
)

func TestPingReqMessageEncodeDecode(t *testing.T) {
	c := message.NewPingReq()
	buf, err := c.Encode()
	assert.NoError(t, err)

	f, _, err := message.ReceiveFrame(bytes.NewReader(buf))
	assert.NoError(t, err)
	assert.Exactly(t, message.PINGREQ, f.Type)
	assert.False(t, f.DUP)
	assert.Equal(t, message.QoS0, f.QoS)
	assert.False(t, f.RETAIN)
}
