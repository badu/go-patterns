package messaging

import (
	"sync"
)

// identifies events topic.
type EventID string

// must be implemented by anything that can be published
type Event interface {
	EventID() EventID
}

// is function that can be subscribed to the event
type EventHandler func(event Event)

// represents active event subscription
type Subscription struct {
	eventID EventID
	id      uint64
}

// allows to subscribe/unsubscribe own event handlers
type Subscriber interface {
	Subscribe(eventID EventID, cb EventHandler) Subscription
	Unsubscribe(id Subscription)
}

// allows to publish own events
type Publisher interface {
	Publish(event Event)
}

// allows to subscribe/unsubscribe to external events and publish own events
type Bus interface {
	Subscriber
	Publisher
}

// returns new event bus
func NewEventBus() Bus {
	b := &bus{
		infos: make(map[EventID]subscriptionInfoList),
	}
	return b
}

type subscriptionInfo struct {
	id uint64
	cb EventHandler
}

type subscriptionInfoList []*subscriptionInfo

type bus struct {
	lock   sync.Mutex
	nextID uint64
	infos  map[EventID]subscriptionInfoList
}

func (b *bus) Subscribe(eventID EventID, cb EventHandler) Subscription {
	b.lock.Lock()
	defer b.lock.Unlock()
	id := b.nextID
	b.nextID++
	sub := &subscriptionInfo{
		id: id,
		cb: cb,
	}
	b.infos[eventID] = append(b.infos[eventID], sub)
	return Subscription{
		eventID: eventID,
		id:      id,
	}
}

func (b *bus) Unsubscribe(subscription Subscription) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if infos, ok := b.infos[subscription.eventID]; ok {
		for idx, info := range infos {
			if info.id == subscription.id {
				infos = append(infos[:idx], infos[idx+1:]...)
				break
			}
		}
		if len(infos) == 0 {
			delete(b.infos, subscription.eventID)
		} else {
			b.infos[subscription.eventID] = infos
		}
	}
}

func (b *bus) Publish(event Event) {
	infos := b.copySubscriptions(event.EventID())
	for _, sub := range infos {
		sub.cb(event)
	}
}

func (b *bus) copySubscriptions(eventID EventID) subscriptionInfoList {
	// External code may subscribe/unsubscribe during iteration over callbacks,
	//  so we need to copy subscribers to invoke callbacks.

	b.lock.Lock()
	defer b.lock.Unlock()
	if infos, ok := b.infos[eventID]; ok {
		return infos
	}
	return subscriptionInfoList{}
}
