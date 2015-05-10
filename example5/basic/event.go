package basic

import (
  "github.com/veandco/go-sdl2/sdl"
)

type EventListener func (event *Event)

type Event struct {
  Sdlevent sdl.Event
  Processed bool
}

// keeps a collection of listeners and can notify them 
type EventEmitter struct {
  listeners []EventListener
}

// notifyAll sends the event to all listeners.
// the listeners can set processed to true, but every 
// subsequent listener will still be notified.
func (self *EventEmitter) NotifyAll(event *Event) {
  for _,listener := range self.listeners {
    listener(event)
  }
}

func (self *EventEmitter) AddEventListener(listener EventListener) {
  if self.listeners == nil {
    self.listeners = []EventListener{}
  }
  self.listeners = append(self.listeners,listener)
}

