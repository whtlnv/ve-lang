package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

// Helpers

func emitEachCharThenEOF(broker *eventBroker.EventBroker, src []byte) {
	for _, char := range src {
		broker.Emit("tokenizer:in:scan", char)
	}
	broker.Emit("tokenizer:in:EOF", nil)
}

// Tests

func TestTokenizerListensToScannedCharactersUntilEOF(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("hello")
	tokenizer := NewTokenizer(broker)

	receivedTokens := [][]byte{}
	broker.On(tokenizer.TokenEvent(), func(data interface{}) {
		receivedTokens = append(receivedTokens, data.([]byte))
	})

	emitEachCharThenEOF(broker, src)

	expectedTokens := [][]byte{
		[]byte("hello"),
		[]byte("EOF"),
	}
	assert.Equal(t, expectedTokens, receivedTokens)
}
