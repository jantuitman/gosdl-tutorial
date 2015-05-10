# GOSDL tutorial


For full tutorial texts see: http://blog.tuitman.org/

##How to run

Install SDL2 on your computer. See https://github.com/veandco/go-sdl2 for instructions.
Install Go (golang) on your computer. Next:

    go get github.com/jantuitman/gosdl-tutorial

Then, install the example you want to run:

    cd $GOPATH/src/github.com/jantuitman/gosdl-tutorial/example1
    go install

All other examples install in the same way.


## List of examples

### Working
  
  All examples require Verdana.ttf to be in /Library/Fonts. This is true on the Mac, but will not always be the case.

  - Example1 - window with eventloop, prints some info about received events.
  - Example2 - Button, ContainerView, combining views into multiple views
  - Example3 - print texts. For now: "Hello world!", 
  - Example4 - a text box. 
  - Example5 - Rewrite of example2, separating the modules a little bit further.

### Planned/TODO.

  - hardware accellerated views
  - further library separation like was done in example2, more generic methods.
  - scrollview that lazily draws content
  - image view
  - label view with possibility of multiple line text.
  - load font from right location.
  - container that does some automatic layout.


