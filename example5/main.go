package main

import ( 
  "github.com/veandco/go-sdl2/sdl_ttf"
  "github.com/jantuitman/gosdl-tutorial/example5/basic"
  "github.com/jantuitman/gosdl-tutorial/example5/widgets"
  "fmt"
)


func main() {
  var runloop  = basic.MainLoop{}

  ttf.Init()
  font,err := ttf.OpenFont("/Library/Fonts/Verdana.ttf",16)
  if err != nil {
    panic(err)
  }
  
  view := widgets.CreateContainerView(0,0,320,240) 
  button1 := widgets.CreateButton(20,20,280,20,font,"Button 1")
  view.AddChild(button1)
  view.AddChild(widgets.CreateButton(20,50,280,20,font,"Button 2"))
  view2 := widgets.CreateContainerView(20,80,280,100)
  view2.AddChild(widgets.CreateButton(20,20,200,20,font,"Button 3"))
  view2.AddChild(widgets.CreateButton(20,50,200,20,font,"Button 4"))
  view.AddChild(view2)
  

  runloop.RootControl = view
  
  button1.OnClick(func(event *basic.Event) {
    fmt.Println("We had a click on button1")
    event.Processed = true
  }) 
  
  runloop.OnEvent(func (event *basic.Event) {
    fmt.Println("First your main event handler can handle it....",event.Sdlevent)
  })
  runloop.OnUnprocessedEvent(func (event *basic.Event) {
    fmt.Println("Or you can handle only unprocessed events....",event.Sdlevent)
  })

  runloop.Run()  
  ttf.Quit()
}
