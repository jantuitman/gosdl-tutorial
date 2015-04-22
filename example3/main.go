package main

import (
  "github.com/veandco/go-sdl2/sdl"
  "github.com/jantuitman/go-sdl2/sdl_ttf"
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
  font,err := ttf.OpenFont("/Library/Fonts/Verdana.ttf",24)
  if err != nil {
    panic(err)
  }
  w,h,err := font.SizeUTF8("Hello World!")
  if err != nil {
    panic(err)
  }
  fmt.Printf("Widht, height of hello world = (%d, %d) \n",w,h)

  renderedTxt1, err := font.RenderUTF8_Blended("Hello world!",sdl.Color{R: 0, G: 255, B: 0, A: 0xFF})
  renderedTxt1.Blit(&sdl.Rect{0,0,renderedTxt1.W,renderedTxt1.H},surface,&sdl.Rect{20,20,renderedTxt1.W,renderedTxt1.H})

  renderedTxt1.Free()
  
  renderedTxt2, err := font.RenderUTF8_Blended_Wrapped("This is an experiment with wrapped text, as you can see. It does not wrap very well if the words don't fit, so the usefullness of this func is debatable.",sdl.Color{R: 255, G: 0, B: 0, A: 0xFF}, 280)
  renderedTxt2.Blit(&sdl.Rect{0,0,renderedTxt2.W,renderedTxt2.H},surface,&sdl.Rect{20,80,renderedTxt2.W,renderedTxt2.H})
  renderedTxt2.Free()
  
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

