package basic;

import (
  "github.com/veandco/go-sdl2/sdl"

)

type ControlLike interface {
  ProcessEvent(event *Event)
  BlitIfNeeded(parentSurface * sdl.Surface) bool
  ParentControl() ControlLike 
  SetParent(control ControlLike)
  GetRect() Rect // relative tov parent
  AbsoluteRect() Rect // relative tov window.
}

type Control struct {
  parent ControlLike
  Rect
  Dirty bool
  surface *sdl.Surface
}


func (self *Control) ParentControl() ControlLike {
  return self.parent
}

func (self *Control) SetParent(c ControlLike) {
  self.parent = c
  self.Dirty = true
}

func (self *Control) GetRect() Rect {
  return self.Rect
}


func (self *Control) AbsoluteRect() Rect {
  if self.parent == nil {
    return self.Rect
  } else {
    result := self.GetRect().offset(self.parent.AbsoluteRect())
    return result
  }
}

func (self *Control) GetSurface() *sdl.Surface {
  if self.surface == nil {
    var err error
    self.surface,err = sdl.CreateRGBSurface(0,int32(self.W),int32(self.H),32,0,0,0,0)
    if err != nil {
      panic(err)
    }
  }
  return self.surface
}


