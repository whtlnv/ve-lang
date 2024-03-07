package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

// Tests

func TestPositionCounterEmitsColumnNumber(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("Hi!")
	counter := NewPositionCounter(broker)

	receivedCharacters := []CharacterEvent{}
	broker.On(counter.CharacterEvent(), func(data interface{}) {
		receivedCharacters = append(receivedCharacters, data.(CharacterEvent))
	})

	emitEachChar(broker, src, counter.ScanEvent())

	expectedCharacters := []CharacterEvent{
		{character: 'H', lineNumber: 1, columnNumber: 1},
		{character: 'i', lineNumber: 1, columnNumber: 2},
		{character: '!', lineNumber: 1, columnNumber: 3},
	}
	assert.Equal(t, expectedCharacters, receivedCharacters)
}

func TestPositionCounterEmitsLineNumber(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("Hi\nHo")
	counter := NewPositionCounter(broker)

	receivedCharacters := []CharacterEvent{}
	broker.On(counter.CharacterEvent(), func(data interface{}) {
		receivedCharacters = append(receivedCharacters, data.(CharacterEvent))
	})

	emitEachChar(broker, src, counter.ScanEvent())

	expectedCharacters := []CharacterEvent{
		{character: 'H', lineNumber: 1, columnNumber: 1},
		{character: 'i', lineNumber: 1, columnNumber: 2},
		{character: '\n', lineNumber: 1, columnNumber: 3},
		{character: 'H', lineNumber: 2, columnNumber: 1},
		{character: 'o', lineNumber: 2, columnNumber: 2},
	}
	assert.Equal(t, expectedCharacters, receivedCharacters)
}

func TestPositionCounterEmitsEOF(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	counter := NewPositionCounter(broker)

	receivedCharacters := []CharacterEvent{}
	broker.On(counter.CharacterEvent(), func(data interface{}) {
		receivedCharacters = append(receivedCharacters, data.(CharacterEvent))
	})

	broker.Emit(counter.EOFEvent(), nil)

	expectedCharacters := []CharacterEvent{
		{character: 4, lineNumber: 1, columnNumber: 1},
	}
	assert.Equal(t, expectedCharacters, receivedCharacters)
}
