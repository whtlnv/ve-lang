package eventBroker

func NewMessageForwarder(broker *EventBroker, inStream string, outStream string) {
	broker.On(inStream, func(data interface{}) {
		broker.Emit(outStream, data)
	})
}
