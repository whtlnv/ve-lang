package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

func TestLexerEmitsEOF(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("")
	lexer := NewLexer(broker)

	eventReceived := false
	broker.On(lexer.EOFEvent(), func(data interface{}) {
		eventReceived = true
	})

	broker.Emit(lexer.ScanEvent(), src)
	assert.True(t, eventReceived)
}

func TestLexerEmitsEachCharacter(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("hello")
	lexer := NewLexer(broker)

	receivedChars := []byte{}
	broker.On(lexer.NewCharacterEvent(), func(data interface{}) {
		receivedChars = append(receivedChars, data.(byte))
	})

	broker.Emit(lexer.ScanEvent(), src)
	assert.Equal(t, src, receivedChars)
}
