package lexer

import "github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"

type Token struct {
	kind  string
	value string
}

type Tokenizer struct {
	broker       *eventBroker.EventBroker
	currentToken []byte
}

// factory

func NewTokenizer(broker *eventBroker.EventBroker) *Tokenizer {
	tokenizer := &Tokenizer{
		broker: broker,
	}

	broker.On(tokenizer.NewCharacterEvent(), func(data interface{}) {
		tokenizer.considerNewCharater(data.(byte))
	})

	broker.On(tokenizer.TokenEndEvent(), func(data interface{}) {
		tokenizer.evaluateToken()
	})

	return tokenizer
}

// do stuff

func (tokenizer *Tokenizer) considerNewCharater(character byte) {
	tokenizer.currentToken = append(tokenizer.currentToken, character)
}

func (tokenizer *Tokenizer) evaluateToken() {
	token := &Token{
		kind:  "PUBLISH",
		value: string(tokenizer.currentToken),
	}

	tokenizer.broker.Emit(tokenizer.NewTokenEvent(), token)
	tokenizer.currentToken = []byte{}
}

// events

func (tokenizer *Tokenizer) NewCharacterEvent() string {
	return "tokenizer:in:character"
}

func (tokenizer *Tokenizer) TokenEndEvent() string {
	return "tokenizer:in:tokenEnd"
}

func (tokenizer *Tokenizer) NewTokenEvent() string {
	return "tokenizer:out:token"
}
