package lexer

import (
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

const /* scan modes */ (
	IDENTIFIER = iota
	STRING
	LINE_COMMENT
)

type Tokenizer struct {
	broker       *eventBroker.EventBroker
	currentToken []byte
	scanMode     int
}

// Clean up later

type Classifier interface {
}

type IdentifierClassifier struct{}
type NumberClassifier struct{}
type OperatorClassifier struct{}
type StringClassifier struct{}
type LineCommentClassifier struct{}
type BlockCommentClassifier struct{}

// constants

func (tokenizer *Tokenizer) Separators() []byte {
	if tokenizer.scanMode == IDENTIFIER {
		return []byte(" .,;()[]{}")
	} else if tokenizer.scanMode == STRING {
		return []byte{}
	} else if tokenizer.scanMode == LINE_COMMENT {
		return []byte("\n")
	} else {
		panic("Unknown scan mode")
	}
}

// factory

func NewTokenizer(broker *eventBroker.EventBroker) *Tokenizer {
	tokenizer := &Tokenizer{
		broker: broker,
	}

	tokenizer.clearCurrentToken()

	broker.On(tokenizer.ScanEvent(), func(data interface{}) {
		tokenizer.processInput(data.(byte))
	})

	broker.On(tokenizer.EOFEvent(), func(data interface{}) {
		// TODO: just evict the token
		tokenizer.processTokenBreak([]byte("EOF"))
	})

	return tokenizer
}

// do stuff

func (tokenizer *Tokenizer) isSeparator(input byte) bool {
	for _, separator := range tokenizer.Separators() {
		if input == separator {
			return true
		}
	}

	return false
}

func (tokenizer *Tokenizer) setScanMode(input byte) {
	if tokenizer.scanMode == IDENTIFIER {
		if input == '"' {
			tokenizer.scanMode = STRING
		}

		if input == '#' {
			tokenizer.scanMode = LINE_COMMENT
		}
	} else if tokenizer.scanMode == STRING {
		if input == '"' {
			tokenizer.scanMode = IDENTIFIER
		}
	}
}

func (tokenizer *Tokenizer) processInput(input byte) {
	tokenizer.setScanMode(input)
	if tokenizer.isSeparator(input) {
		tokenizer.processTokenBreak([]byte{input})
	} else {
		tokenizer.currentToken = append(tokenizer.currentToken, input)
	}
}

func (tokenizer *Tokenizer) evictToken() {
	tokenizer.broker.Emit(tokenizer.TokenEvent(), tokenizer.currentToken)
	tokenizer.clearCurrentToken()
}

func (tokenizer *Tokenizer) processTokenBreak(tokenBreak []byte) {
	tokenizer.evictToken()
	tokenizer.broker.Emit(tokenizer.TokenEvent(), tokenBreak)
}

func (tokenizer *Tokenizer) clearCurrentToken() {
	tokenizer.currentToken = []byte{}
	tokenizer.scanMode = IDENTIFIER
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
