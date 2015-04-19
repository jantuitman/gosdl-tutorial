package main

import (
  "github.com/veandco/go-sdl2/sdl"
  "github.com/veandco/go-sdl2/sdl_ttf"
  "fmt"
)

func main() {
    sdl.Init(sdl.INIT_EVERYTHING)

    window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        320, 480, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()
    test(window)
    eventLoop()
}


func test(window *sdl.Window) {
  surface, err := window.GetSurface()
  if err != nil {
      panic(err)
  }


  ttf.Init()
  font,err := ttf.OpenFont("/Library/Fonts/Verdana.ttf",12)
  if err != nil {
    panic(err)
  }
  var renderedTxt *sdl.Surface = font.RenderText_Shaded("Hello world!",sdl.Color{R: 0, G: 0, B: 0, A: 0xFF},sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF})
  renderedTxt.Blit(&sdl.Rect{0,0,renderedTxt.W,renderedTxt.H},surface,&sdl.Rect{50,50,renderedTxt.W,renderedTxt.H})
  window.UpdateSurface()
  ttf.Quit()
}


func eventLoop() {
  quiting := false
  for !quiting {
    event := sdl.PollEvent()
    switch t := event.(type) {
      case *sdl.QuitEvent:
          fmt.Println("QuitEvent",event,t)
          quiting = true
          break
      case *sdl.KeyDownEvent:
          fmt.Println("KeyboardEvent",event,t)
          break
      default:
          fmt.Printf("UNHANDLED event of type %T\n",t)
          break
    }
    sdl.Delay(100)
  }
  sdl.Quit()
}

