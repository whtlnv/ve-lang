package eventBroker

type EventBroker struct {
	listeners map[string][]func(interface{})
}

func NewEventBroker() *EventBroker {
	return &EventBroker{
		listeners: make(map[string][]func(interface{})),
	}
}

func (e *EventBroker) On(event string, listener func(interface{})) {
	e.listeners[event] = append(e.listeners[event], listener)
}

func (e *EventBroker) Emit(event string, data interface{}) {
	for _, listener := range e.listeners[event] {
		listener(data)
	}
}
