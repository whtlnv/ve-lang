package lexer

import (
	"github.com/whtlnv/ve-lang/go-bootstrap/pkg/eventBroker"
)

type PositionCounter struct {
	broker       *eventBroker.EventBroker
	lineNumber   int
	columnNumber int
}

type CharacterEvent struct {
	character    byte
	lineNumber   int
	columnNumber int
}

// factory

func NewPositionCounter(broker *eventBroker.EventBroker) *PositionCounter {
	positionCounter := &PositionCounter{
		broker:       broker,
		lineNumber:   1,
		columnNumber: 0,
	}

	broker.On(positionCounter.ScanEvent(), func(data interface{}) {
		positionCounter.processInput(data.(byte))
	})

	broker.On(positionCounter.EOFEvent(), func(data interface{}) {
		eofAsByte := byte(4)
		positionCounter.processInput(eofAsByte)
	})

	return positionCounter
}

// do stuff

func (counter *PositionCounter) processInput(char byte) {
	counter.columnNumber++

	counter.broker.Emit(counter.CharacterEvent(), CharacterEvent{
		character:    char,
		lineNumber:   counter.lineNumber,
		columnNumber: counter.columnNumber,
	})

	if char == '\n' {
		counter.lineNumber++
		counter.columnNumber = 0
	}
}

// events

func (counter *PositionCounter) ScanEvent() string {
	return "position:in:scan"
}

func (counter *PositionCounter) EOFEvent() string {
	return "position:in:EOF"
}

func (counter *PositionCounter) CharacterEvent() string {
	return "position:out:character"
}
