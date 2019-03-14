package main

import (
	"github.com/zasy/observer"
	"time"
)

func main() {
	// Initialize a new Notifier
	n := observer.EventNotifier{
		Observers: map[observer.Observer]struct{}{},
	}

	// Register a couple of observes.
	n.Register(&observer.EventObserver{Id: 1})
	n.Register(&observer.EventObserver{Id: 2})

	// A simple loop publishing the current Unix timestamp to observers.
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <- stop:
			return
		case t := <-tick:
			n.Notify(observer.Event{Data: t.UnixNano()})
		}
	}
}