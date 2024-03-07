package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

// Helpers

func emitEachChar(broker *eventBroker.EventBroker, src []byte, topic string) {
	for _, char := range src {
		broker.Emit(topic, char)
	}
}

// Tests

func TestTokenizerGroupsCharactersUntilEOF(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("hello")
	tokenizer := NewTokenizer(broker)

	receivedTokens := [][]byte{}
	broker.On(tokenizer.TokenEvent(), func(data interface{}) {
		receivedTokens = append(receivedTokens, data.([]byte))
	})

	emitEachChar(broker, src, tokenizer.ScanEvent())
	broker.Emit(tokenizer.EOFEvent(), nil)

	expectedTokens := [][]byte{
		[]byte("hello"),
		[]byte("EOF"),
	}
	assert.Equal(t, expectedTokens, receivedTokens)
}

func tokenizerGroupsCharactersUntilBreak(t *testing.T, separator byte) {
	broker := eventBroker.NewEventBroker()

	src := []byte("hello" + string(separator) + "world")
	tokenizer := NewTokenizer(broker)

	receivedTokens := [][]byte{}
	broker.On(tokenizer.TokenEvent(), func(data interface{}) {
		receivedTokens = append(receivedTokens, data.([]byte))
	})

	emitEachChar(broker, src, tokenizer.ScanEvent())
	broker.Emit(tokenizer.EOFEvent(), nil)

	expectedTokens := [][]byte{
		[]byte("hello"),
		{separator},
		[]byte("world"),
		[]byte("EOF"),
	}
	assert.Equal(t, expectedTokens, receivedTokens)
}

func TestTokenizerIdentifierBreaks(t *testing.T) {
	separators := []byte(" .,;()[]{}")

	for _, separator := range separators {
		tokenizerGroupsCharactersUntilBreak(t, separator)
	}
}

func TestTokenizerGroupsCharactersInAString(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("\"hello world\" \"foo bar\"")
	tokenizer := NewTokenizer(broker)

	receivedTokens := [][]byte{}
	broker.On(tokenizer.TokenEvent(), func(data interface{}) {
		receivedTokens = append(receivedTokens, data.([]byte))
	})

	emitEachChar(broker, src, tokenizer.ScanEvent())
	broker.Emit(tokenizer.EOFEvent(), nil)

	expectedTokens := [][]byte{
		[]byte("\"hello world\""),
		[]byte(" "),
		[]byte("\"foo bar\""),
		[]byte("EOF"),
	}
	assert.Equal(t, expectedTokens, receivedTokens)
}

func TestTokenizerGroupsCharactersInALineComment(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte("# hello world")
	tokenizer := NewTokenizer(broker)

	receivedTokens := [][]byte{}
	broker.On(tokenizer.TokenEvent(), func(data interface{}) {
		receivedTokens = append(receivedTokens, data.([]byte))
	})

	emitEachChar(broker, src, tokenizer.ScanEvent())
	broker.Emit(tokenizer.EOFEvent(), nil)

	expectedTokens := [][]byte{
		[]byte("# hello world"),
		[]byte("EOF"),
	}
	assert.Equal(t, expectedTokens, receivedTokens)
}

// func TestTokenizerEmitsDifferentKindOfTokens(t *testing.T) {
// 	broker := eventBroker.NewEventBroker()

// 	src := []byte("\"hello world\"# this is a string")
// 	tokenizer := NewTokenizer(broker)

// 	receivedTokens := [][]byte{}
// 	broker.On(tokenizer.TokenEvent(), func(data interface{}) {
// 		receivedTokens = append(receivedTokens, data.([]byte))
// 	})

// 	emitEachCharThenEOF(broker, src)

// 	expectedTokens := [][]byte{
// 		[]byte("\"hello world\""),
// 		[]byte("# this is a string"),
// 		[]byte("EOF"),
// 	}
// 	assert.Equal(t, expectedTokens, receivedTokens)
// }

// func TestTokenizerGroupsOperandCharacters(t *testing.T) {
// 	broker := eventBroker.NewEventBroker()

// 	src := []byte("1+2")
// 	tokenizer := NewTokenizer(broker)

// 	receivedTokens := [][]byte{}
// 	broker.On(tokenizer.TokenEvent(), func(data interface{}) {
// 		receivedTokens = append(receivedTokens, data.([]byte))
// 	})

// 	emitEachCharThenEOF(broker, src)

// 	expectedTokens := [][]byte{
// 		[]byte("1"),
// 		[]byte("+"),
// 		[]byte("2"),
// 		[]byte("EOF"),
// 	}
// 	assert.Equal(t, expectedTokens, receivedTokens)
// }
