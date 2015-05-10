package widgets

import (
  "github.com/veandco/go-sdl2/sdl"
  "github.com/jantuitman/gosdl-tutorial/example5/basic"

  //"github.com/veandco/go-sdl2/sdl_ttf"
  //"fmt"
)

type ContainerView struct {
  basic.Control
  children []basic.ControlLike
  surface *sdl.Surface
}

func CreateContainerView(x int,y int,w int,h int) *ContainerView {
  var result = ContainerView{}
  result.X = x
  result.Y = y
  result.W = w
  result.H = h
  return &result
}

func (self *ContainerView) AddChild(b basic.ControlLike) {
  b.SetParent(self)
  self.children = append(self.children,b)
  self.Dirty = true
}

func (self *ContainerView) ProcessEvent(event *basic.Event) {
  for _,control := range self.children {
    if event.Processed {
      break 
    }
    control.ProcessEvent(event)
  } 
}

func (self *ContainerView) BlitIfNeeded(parentSurface *sdl.Surface) bool {
  result := false
  for _,control := range self.children {
    self.Dirty = self.Dirty || control.BlitIfNeeded(self.GetSurface())
  }
  if self.Dirty {
    var d = self.GetRect()
    self.GetSurface().Blit(
      &sdl.Rect{0,0,int32(self.W),int32(self.H)},
      parentSurface,
      &sdl.Rect{int32(d.X),int32(d.Y),int32(d.W),int32(d.H)})
    self.Dirty = false
    result = true
  }
  return result
}