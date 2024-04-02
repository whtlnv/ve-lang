package classifiers

import (
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

type StringClassifier struct {
	broker       *eventBroker.EventBroker
	currentToken []byte
}

// factory

func NewStringClassifier(broker *eventBroker.EventBroker) *StringClassifier {
	classifier := &StringClassifier{
		broker: broker,
	}

	classifier.clearCurrentToken()

	broker.On(classifier.ScanEvent(), func(data interface{}) {
		classifier.processInput(data.(byte))
	})

	return classifier
}

// do stuff

func (classifier *StringClassifier) processInput(char byte) {}

func (classifier *StringClassifier) clearCurrentToken() {}

// events

func (classifier *StringClassifier) ScanEvent() string {
	return "string:in:scan"
}

func (classifier *StringClassifier) TokenEvent() string {
	return "string:out:token"
}
