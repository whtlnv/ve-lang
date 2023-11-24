package eventBroker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInEventsShouldEmitOutEvents(t *testing.T) {
	broker := NewEventBroker()

	inStream := "input:message"
	outStream := "output:publish"

	eventReceived := false
	broker.On(outStream, func(data interface{}) {
		eventReceived = true
	})

	NewMessageForwarder(broker, inStream, outStream)

	broker.Emit(inStream, "test")
	assert.True(t, eventReceived)
}
