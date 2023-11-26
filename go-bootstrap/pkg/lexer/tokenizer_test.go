package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

func TestTokenizerCanTokenizePublishKeyword(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	tokenizer := NewTokenizer(broker)

	want := &Token{
		kind:  "PUBLISH",
		value: "publish",
	}

	var got *Token
	broker.On(tokenizer.NewTokenEvent(), func(data interface{}) {
		got = data.(*Token)
	})

	source := []byte("publish")
	for _, char := range source {
		broker.Emit(tokenizer.NewCharacterEvent(), char)
	}
	broker.Emit(tokenizer.TokenEndEvent(), nil)

	assert.Equal(t, want, got)
}
