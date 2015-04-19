package main

import (
  "github.com/veandco/go-sdl2/sdl"
  "fmt"
)

func main() {
    sdl.Init(sdl.INIT_EVERYTHING)

    window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        800, 600, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()
    test1(window)
}


/* draws a rect. and listens for keys.*/
func test1(window *sdl.Window) {
  surface, err := window.GetSurface()
  if err != nil {
      panic(err)
  }

  rect := sdl.Rect{0, 0, 200, 200}
  surface.FillRect(&rect, 0xffff0000)
  window.UpdateSurface()
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