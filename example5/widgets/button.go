package widgets;

import (
  "github.com/veandco/go-sdl2/sdl"
  "github.com/veandco/go-sdl2/sdl_ttf"
  "github.com/jantuitman/gosdl-tutorial/example5/basic"

  "fmt"
)

type Button struct {
  // private properties
  basic.Control
  onClick basic.EventEmitter
  font *ttf.Font
  text string
 }

func CreateButton(x int,y int,w int, h int,font *ttf.Font, text string) *Button {
  var result Button = Button{}
  result.X = x
  result.Y = y
  result.W = w
  result.H = h
  result.font = font
  result.text = text
  result.Dirty = true 
  return &result
}

func (self *Button) OnClick(eventListener basic.EventListener) {
  self.onClick.AddEventListener(eventListener)
}



func (self *Button) ProcessEvent(event *basic.Event) {
  switch event.Sdlevent.(type) {
    case *sdl.MouseButtonEvent:
      e := event.Sdlevent.(*sdl.MouseButtonEvent)
      var d = self.AbsoluteRect()
      fmt.Printf("in button %s dimensions are: %d,%d,%d,%d \n",self.text,d.X,d.Y,d.W,d.H)
      if int32(d.X) <= e.X && e.X <= int32(d.X + d.W) {
        if int32(d.Y) <= e.Y && e.Y <= int32(d.Y + d.H) {
          self.onClick.NotifyAll(event)
        }
      }
  }
}


func (self *Button) BlitIfNeeded(parentSurface *sdl.Surface) bool {
  result := false
  if self.Dirty {
    // fill a rectangle
    self.GetSurface().FillRect(&sdl.Rect{0,0,int32(self.W),int32(self.H)},0xFFFFFFFF)
    // blit the text
    if self.text != "" {
      var txt *sdl.Surface
      txt,err := self.font.RenderUTF8_Blended(self.text,sdl.Color{R: 0, G: 0, B: 0, A: 0xFF})
      if err != nil {
        panic(err)
      }
      txt.Blit(&sdl.Rect{int32(0),0,txt.W - int32(0),txt.H}, self.GetSurface(), &sdl.Rect{2,0,txt.W,txt.H})   
      txt.Free()
    }

    // blit it
    var rect = self.GetRect()
    self.GetSurface().Blit(
      &sdl.Rect{0,0,int32(self.W),int32(self.H)},
      parentSurface,
      &sdl.Rect{int32(rect.X),int32(rect.Y),int32(rect.W),int32(rect.H)})
    self.Dirty = false 
    result = true
  }
  return result 
}
