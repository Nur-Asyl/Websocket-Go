package ws

type Observer interface {
	Update(message *Message)
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(message *Message)
}
