package classifiers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whtlnv/ve-lang/go-bootstrap/internal/testUtilities"
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

func TestStringClassifierFindsAString(t *testing.T) {
	broker := eventBroker.NewEventBroker()

	src := []byte(`"hello"`)
	classifier := NewStringClassifier(broker)

	receivedTokens := [][]byte{}
	broker.On(classifier.TokenEvent(), func(data interface{}) {
		receivedTokens = append(receivedTokens, data.([]byte))
	})

	testUtilities.EmitEachChar(broker, src, classifier.ScanEvent())

	expectedTokens := [][]byte{
		[]byte("hello"),
	}
	assert.Equal(t, expectedTokens, receivedTokens)
}
