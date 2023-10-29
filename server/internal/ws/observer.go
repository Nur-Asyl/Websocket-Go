package ws

type Subscriber interface {
	Update(message *Message)
}

type ConcreteSubscriber interface {
	Subscribe(sub Subscriber)
	Unsubscribe(sub Subscriber)
	Notify(message *Message)
}
