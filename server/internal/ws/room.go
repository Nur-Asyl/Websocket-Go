package ws

type Room struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Clients     map[string]*Client `json:"clients"`
	Subscribers []Subscriber
}

func (r *Room) Subscribe(observer Subscriber) {
	r.Subscribers = append(r.Subscribers, observer)
}

func (r *Room) Unsubscribe(observer Subscriber) {
	for i, obs := range r.Subscribers {
		if obs == observer {
			r.Subscribers = append(r.Subscribers[:i], r.Subscribers[i+1:]...)
			break
		}
	}
}

func (r *Room) Notify(message *Message) {
	for _, observer := range r.Subscribers {
		observer.Update(message)
	}
}
