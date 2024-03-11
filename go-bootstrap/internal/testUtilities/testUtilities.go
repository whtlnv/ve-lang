package testUtilities

import (
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

func EmitEachChar(broker *eventBroker.EventBroker, src []byte, topic string) {
	for _, char := range src {
		broker.Emit(topic, char)
	}
}
