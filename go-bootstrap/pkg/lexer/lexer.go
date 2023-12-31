package lexer

import "github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"

type Lexer struct {
	broker *eventBroker.EventBroker
	source []byte
}

// factory

func NewLexer(broker *eventBroker.EventBroker) *Lexer {
	lexer := &Lexer{
		broker: broker,
	}

	broker.On(lexer.ScanEvent(), func(data interface{}) {
		lexer.source = data.([]byte)
		lexer.scanSource()
	})

	return lexer
}

// do stuff

func (lexer *Lexer) scanSource() {
	for _, char := range lexer.source {
		lexer.broker.Emit(lexer.NewCharacterEvent(), char)
	}

	lexer.broker.Emit(lexer.EOFEvent(), nil)
}

// events

func (lexer *Lexer) ScanEvent() string {
	return "lexer:in:scan"
}

func (lexer *Lexer) NewCharacterEvent() string {
	return "lexer:out:didScan"
}

func (lexer *Lexer) EOFEvent() string {
	return "lexer:out:didFinishScan"
}
