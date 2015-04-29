package main

import (
  "github.com/veandco/go-sdl2/sdl"
  "github.com/veandco/go-sdl2/sdl_ttf"
  "fmt"
)

type ControlEvent struct {
  sdlevent sdl.Event
  processed bool 
}


type BasicControl interface {
  ProcessEvent(event *ControlEvent)
  BlitIfNeeded(parentSurface * sdl.Surface) bool
  ParentControl() BasicControl 
  setParent(control BasicControl)
  getDimensions() Dimensions // relative tov parent
  absoluteDimensions() Dimensions // relative tov window.
  isDirty() bool
}


type Dimensions struct {
  x int
  y int
  w int
  h int
}

func (d Dimensions) offset(other Dimensions) Dimensions {
  return Dimensions{other.x + d.x,other.y + d.y, d.w,d.h}
}

type Basic struct {
  parent BasicControl
  Dimensions
  dirty bool
  surface *sdl.Surface
}

func (b *Basic) isDirty() bool {
  return b.dirty 
}

func (b *Basic) getSurface() *sdl.Surface {
  if b.surface == nil {
    var err error
    b.surface,err = sdl.CreateRGBSurface(0,int32(b.w),int32(b.h),32,0,0,0,0)
    fmt.Println("surface",b.surface)
    if err != nil {
      panic(err)
    }
  }
  return b.surface
}

func (b *Basic) ParentControl() BasicControl {
  return b.parent
}
func (b *Basic) setParent(parent BasicControl)  {
  b.parent = parent
  b.dirty = true
}

func (b *Basic) absoluteDimensions() Dimensions {
  if b.parent == nil {
    return b.getDimensions()
  } else {
    return b.Dimensions.offset(b.ParentControl().absoluteDimensions())
  }
}

func (b *Basic) getDimensions() Dimensions {
  return b.Dimensions
}

type ContainerView struct {
  Basic
  children []BasicControl
  surface *sdl.Surface
}

func CreateContainerView(x int,y int,w int,h int) *ContainerView {
  var result = ContainerView{}
  result.x = x
  result.y = y
  result.w = w
  result.h = h
  return &result
}

func (self *ContainerView) addChild(b BasicControl) {
  b.setParent(self)
  self.children = append(self.children,b)
  self.dirty = true
}

func (self *ContainerView) ProcessEvent(event *ControlEvent) {
  for _,control := range self.children {
    if event.processed {
      break 
    }
    control.ProcessEvent(event)
  } 
}

func (self *ContainerView) BlitIfNeeded(parentSurface *sdl.Surface) bool {
  result := false
  for _,control := range self.children {
    self.dirty = self.dirty || control.BlitIfNeeded(self.getSurface())
  }
  if self.dirty {
    var d = self.getDimensions()
    self.getSurface().Blit(
      &sdl.Rect{0,0,int32(self.w),int32(self.h)},
      parentSurface,
      &sdl.Rect{int32(d.x),int32(d.y),int32(d.w),int32(d.h)})
    self.dirty = false
    result = true
  }
  return result
}

type Button struct {
  // private properties
  Basic
  font *ttf.Font
  text string
 }

func CreateButton(x int,y int,w int, h int,font *ttf.Font, text string) *Button {
  var result Button = Button{}
  result.x = x
  result.y = y
  result.w = w
  result.h = h
  result.font = font
  result.text = text
  result.dirty = true 
  return &result
} 

func (self *Button) ProcessEvent(event *ControlEvent) {
  switch event.sdlevent.(type) {
    case *sdl.MouseButtonEvent:
      e := event.sdlevent.(*sdl.MouseButtonEvent)
      var d = self.absoluteDimensions()
      fmt.Printf("in button %s dimensions are: %d,%d,%d,%d \n",self.text,d.x,d.y,d.w,d.h)
      if int32(d.x) <= e.X && e.X <= int32(d.x + d.w) {
        if int32(d.y) <= e.Y && e.Y <= int32(d.y + d.h) {
          fmt.Println("Button was pressed:",self.text)
          event.processed = true
        }
      }
  }
}


func (self *Button) BlitIfNeeded(parentSurface *sdl.Surface) bool {
  result := false
  if self.dirty {
    fmt.Println("Blitting the button")
    // fill a rectangle
    self.getSurface().FillRect(&sdl.Rect{0,0,int32(self.w),int32(self.h)},0xFFFFFFFF)
    // blit the text
    if self.text != "" {
      var txt *sdl.Surface
      txt,err := self.font.RenderUTF8_Blended(self.text,sdl.Color{R: 0, G: 0, B: 0, A: 0xFF})
      if err != nil {
        panic(err)
      }
      txt.Blit(&sdl.Rect{int32(0),0,txt.W - int32(0),txt.H}, self.surface, &sdl.Rect{2,0,txt.W,txt.H})   
      txt.Free()
    }

    // blit it
    var d = self.getDimensions()
    self.getSurface().Blit(
      &sdl.Rect{0,0,int32(self.w),int32(self.h)},
      parentSurface,
      &sdl.Rect{int32(d.x),int32(d.y),int32(d.w),int32(d.h)})
    self.dirty = false 
    fmt.Println("dirty now is",self.dirty)
    result = true
  }
  return result 
}






func main() {
    sdl.Init(sdl.INIT_EVERYTHING)

    window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        320, 480, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()
    test(window)
}


func test(window *sdl.Window) {
  surface, err := window.GetSurface()
  if err != nil {
      panic(err)
  }


  ttf.Init()
  font,err := ttf.OpenFont("/Library/Fonts/Verdana.ttf",16)
  if err != nil {
    panic(err)
  }
  w,h,err := font.SizeUTF8("Hello World!")
  if err != nil {
    panic(err)
  }
  fmt.Printf("Widht, height of hello world = (%d, %d) \n",w,h)

  container := CreateContainerView(0,0,320,480)
  container.addChild(CreateButton(20,20,280,20,font,"Button 1"))
  container.addChild(CreateButton(20,50,280,20,font,"Button 2"))
  container2 := CreateContainerView(20,80,280,100)
  container2.addChild(CreateButton(20,20,260,20,font,"Button 3"))
  container2.addChild(CreateButton(20,50,260,20,font,"Button 4"))
  container.addChild(container2)


  

  quiting := false
  for !quiting {
    event := sdl.PollEvent()
    switch event.(type) {
      case *sdl.QuitEvent:
          fmt.Println("QuitEvent",event)
          quiting = true
          break
      case nil:
        // don't process anything.
      default:
          container.ProcessEvent(&ControlEvent{event,false})
    }
    if container.BlitIfNeeded(surface) {
      window.UpdateSurface()
    }

    
  }

  ttf.Quit()
  sdl.Quit()
}

