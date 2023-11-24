package eventBroker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventBrokerEmitsEvents(t *testing.T) {
	broker := NewEventBroker()

	event := "test"
	eventReceived := false

	broker.On(event, func(data interface{}) {
		eventReceived = true
	})

	broker.Emit(event, nil)
	assert.True(t, eventReceived)
}

func TestEventBrokerEmitsEventData(t *testing.T) {
	broker := NewEventBroker()

	event := "test"
	want := "test data"
	got := ""

	broker.On(event, func(data interface{}) {
		got = data.(string)
	})

	broker.Emit(event, want)
	assert.Equal(t, want, got)
}
