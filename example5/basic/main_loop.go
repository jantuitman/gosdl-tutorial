package basic

import (
  "github.com/veandco/go-sdl2/sdl"

)

type MainLoop struct {
  RootControl ControlLike
  onEvent EventEmitter
  onUnprocessedEvent EventEmitter

}




func (self *MainLoop) OnEvent(eventListener EventListener) {
  self.onEvent.AddEventListener(eventListener)
}

func (self *MainLoop) OnUnprocessedEvent(eventListener EventListener) {
  self.onUnprocessedEvent.AddEventListener(eventListener)
}


// handles one event.
func (self *MainLoop) processEvent(event *Event) {
  // first let the listeners do their job
  self.onEvent.NotifyAll(event)

  // next, let the controls do their job.
  if !event.Processed && self.RootControl != nil {
    self.RootControl.ProcessEvent(event)
  }

  // last chance: unprocessed events 
  if !event.Processed {
    self.onUnprocessedEvent.NotifyAll(event)
  }
}

// loops : polls events, draws dirty components.
func (self *MainLoop) Run() {

  window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
      320, 480, sdl.WINDOW_SHOWN)
  if err != nil {
      panic(err)
  }
  // get the surface
  surface, err := window.GetSurface()
  if err != nil {
      panic(err)
  }

  defer window.Destroy()
  
  quiting := false 
  for !quiting {

    // events
    e := sdl.PollEvent()
    switch e.(type) {
    case *sdl.QuitEvent:
      quiting = true
      break
    case nil:
      // don't process anything
    default:
      event := Event{e,false}
      self.processEvent(&event)
    }

    // painting
    if self.RootControl != nil && self.RootControl.BlitIfNeeded(surface) {
      window.UpdateSurface()
    }
  }

  sdl.Quit()
}