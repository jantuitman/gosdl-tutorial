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

  - Example1 - window with eventloop, prints some info about received events.
  - Example3 - print texts. For now: "Hello world!", requires Verdana.ttf to be in /Library/Fonts

### Planned

  - Example2 - hardware accellerated clipping and combining views
  - Example4 - a text box

### Unplanned

More to come soon.