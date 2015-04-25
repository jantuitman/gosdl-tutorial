package main

import (
  "github.com/veandco/go-sdl2/sdl"
  "github.com/jantuitman/go-sdl2/sdl_ttf"
  "fmt"
  "bytes"
)


type TextBox struct {
  // private properties
  x int
  y int
  w int
  h int
  surface *sdl.Surface
  focus bool
  text string
  dirty bool
  font  *ttf.Font
  cursorX int
}

func CreateTextBox(x int,y int,w int, h int,font *ttf.Font) *TextBox {
  var result *TextBox = &TextBox{}
  result.x = x
  result.y = y
  result.w = w
  result.h = h
  result.font = font

  return result
} 

// returns true if it blitted. so that containing controls may blit also .
func (self *TextBox) BlitIfNeeded(parentSurface * sdl.Surface) bool {
  updated := false
  if self.surface == nil {
    self.surface = self.initSurface()
    self.dirty = true
  }
  if self.dirty {
    // draw the background
    self.surface.FillRect(&sdl.Rect{0,0,int32(self.w),int32(self.h)},0xFFFFFFFF)
    
    // find out where the cursor is
    var cursorLocation = 0
    if self.cursorX > 0 {
      w,_,err := self.font.SizeUTF8(self.text[0:self.cursorX])
      if err != nil {
        panic(err)
      }
      cursorLocation = cursorLocation + w + 1
    }

    // check if the cursorlocation is almost outside the window, and if so, make a scrolloffset
    var scrollOffset = 0
    if cursorLocation > self.w - 40 {
      scrollOffset = cursorLocation - ( self.w - 40 )
    }

    // draw the text
    if self.text != "" {
      var txt *sdl.Surface
      txt,err := self.font.RenderUTF8_Blended(self.text,sdl.Color{R: 0, G: 0, B: 0, A: 0xFF})
      if err != nil {
        panic(err)
      }
      txt.Blit(&sdl.Rect{int32(scrollOffset),0,txt.W - int32(scrollOffset),txt.H}, self.surface, &sdl.Rect{2,0,txt.W,txt.H})   
    }
    // blit the cursor
    if self.focus {
      self.surface.FillRect(&sdl.Rect{int32(cursorLocation + 2 - scrollOffset),2,2,int32(self.h - 4)},0x000000FF)
    }

    self.surface.Blit(&sdl.Rect{0,0,int32(self.w),int32(self.h)}, parentSurface, &sdl.Rect{int32(self.x),int32(self.y),int32(self.w),int32(self.h)})
    self.dirty = false
    updated = true
  }
  return updated
}

func (self *TextBox) initSurface() *sdl.Surface {
  surface,err := sdl.CreateRGBSurface(0,int32(self.w),int32(self.h),32,0,0,0,0)
  if err != nil {
    panic(err)
  }
  return surface
}

func (self *TextBox) SetFocus(newValue bool) {
  if newValue != self.focus {
    self.dirty = true
    self.focus = newValue
  }
}

func (self *TextBox) SetText(newValue string) {
  if newValue != self.text {
    self.dirty = true
    self.text = newValue
  }
}

func (self *TextBox) handleKeyDown(keyEvent *sdl.KeyDownEvent) {
  switch keyEvent.Keysym.Sym {
    case sdl.K_TAB:
      self.SetFocus(!self.focus)
    case sdl.K_RETURN:
      self.SetFocus(!self.focus)
    case sdl.K_BACKSPACE:
      if self.cursorX > 0 {
        self.text = self.text[0:self.cursorX - 1] + self.text[self.cursorX:]
        self.cursorX--
        self.dirty = true
      }
    case sdl.K_DELETE:
      if self.cursorX < len(self.text) {
        self.text = self.text[0:self.cursorX] + self.text[self.cursorX+1:]
        self.dirty = true
      }
    case sdl.K_HOME, sdl.K_PAGEUP:       
      if self.cursorX != 0 {
        self.cursorX = 0
        self.dirty = true
      }
    case sdl.K_END, sdl.K_PAGEDOWN:        
      if self.cursorX != len(self.text) {
        self.cursorX = len(self.text)
        self.dirty = true
      }
    case sdl.K_RIGHT:      
      if (self.cursorX < len(self.text)) {
        self.cursorX++
        self.dirty = true
      }
    case sdl.K_LEFT:       
      if self.cursorX > 0 {
        self.cursorX--
        self.dirty = true
      }
    case sdl.K_DOWN:       
      fmt.Println("DOWN pressed")
    case sdl.K_UP:         
      fmt.Println("UP pressed")
    default:
    // no action required  
  }
}

func (self *TextBox) HandleEvent(event sdl.Event) {
  switch event.(type) {
    case *sdl.KeyDownEvent:
      fmt.Println("KeyDownEvent",event)
      keyEvent := event.(*sdl.KeyDownEvent)
      self.handleKeyDown(keyEvent)
    case *sdl.TextInputEvent:
      if self.focus {
        fmt.Println("TextInputEvent while focused",event)
        // text is
        textInput := event.(*sdl.TextInputEvent)  
        n := bytes.Index(textInput.Text[:], []byte{0})
        self.text = self.text[0:self.cursorX] + string(textInput.Text[:n]) + self.text[self.cursorX:]
        self.dirty = true
        self.cursorX++
        fmt.Println("String is nu ",self.text)
      }
    default:
      // No action required  
  }


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

  renderedTxt1, err := font.RenderUTF8_Blended("Textbox demo",sdl.Color{R: 0, G: 255, B: 0, A: 0xFF})
  renderedTxt1.Blit(&sdl.Rect{0,0,renderedTxt1.W,renderedTxt1.H},surface,&sdl.Rect{20,20,renderedTxt1.W,renderedTxt1.H})

  renderedTxt1.Free()

  textbox1 := CreateTextBox(20,50,280,20,font)


  
  renderedTxt2, err := font.RenderUTF8_Blended_Wrapped("Press on the textbox to focus it. After that a cursor should appear. You can then type text (with regular keyboards, iOS/android on screen keyboard not yet supported). Use TAB or ENTER to remove focus ",sdl.Color{R: 255, G: 0, B: 0, A: 0xFF}, 280)
  renderedTxt2.Blit(&sdl.Rect{0,0,renderedTxt2.W,renderedTxt2.H},surface,&sdl.Rect{20,80,renderedTxt2.W,renderedTxt2.H})
  renderedTxt2.Free()
  

  quiting := false
  for !quiting {
    event := sdl.PollEvent()
    switch event.(type) {
      case *sdl.QuitEvent:
          fmt.Println("QuitEvent",event)
          quiting = true
          break
      default:
          textbox1.HandleEvent(event)
          break
    }
    if textbox1.BlitIfNeeded(surface) {
      window.UpdateSurface()
    }

    
  }

  ttf.Quit()
  sdl.Quit()
}

