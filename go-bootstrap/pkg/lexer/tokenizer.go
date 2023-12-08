package lexer

import "github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"

type Tokenizer struct {
	broker       *eventBroker.EventBroker
	currentToken []byte
}

// factory
func NewTokenizer(broker *eventBroker.EventBroker) *Tokenizer {
	tokenizer := &Tokenizer{
		broker:       broker,
		currentToken: []byte{},
	}

	broker.On(tokenizer.ScanEvent(), func(data interface{}) {
		tokenizer.processInput(data.(byte))
	})

	broker.On(tokenizer.EOFEvent(), func(data interface{}) {
		tokenizer.processTokenBreak([]byte("EOF"))
	})

	return tokenizer
}

// do stuff

func (tokenizer *Tokenizer) processInput(input byte) {
	// if input == ' ' {
	// 	tokenizer.processTokenBreak()
	// } else {
	tokenizer.currentToken = append(tokenizer.currentToken, input)
	// }
}

func (tokenizer *Tokenizer) processTokenBreak(tokenBreak []byte) {
	tokenizer.broker.Emit(tokenizer.TokenEvent(), tokenizer.currentToken)
	tokenizer.broker.Emit(tokenizer.TokenEvent(), tokenBreak)
	tokenizer.currentToken = []byte{}
}

// events

func (tokenizer *Tokenizer) ScanEvent() string {
	return "tokenizer:in:scan"
}

func (tokenizer *Tokenizer) EOFEvent() string {
	return "tokenizer:in:EOF"
}

func (tokenizer *Tokenizer) TokenEvent() string {
	return "tokenizer:out:token"
}
