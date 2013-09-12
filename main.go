package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	filename := os.Getenv("HOME") + "/.gosnake.log"
	logfile, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logfile)

	// initialize termbox
	err := termbox.Init()
	if err != nil {
		fmt.Println("Could not start termbox for gosnake.")
		log.Printf("Cannot start gomatrix, termbox.Init() gave an error:\n%s\n", err)
		os.Exit(1)
	}
	termbox.HideCursor()

	var snake = NewSnake()
	var scene = Scene{character: snake}

	// go
	go func() {
		scene.SetSize(termbox.Size())
		for {
			<-time.After(40 * time.Millisecond)
			scene.Draw()
		}
	}()

	// make chan for tembox events and run poller to send events on chan
	eventChan := make(chan termbox.Event)
	go func() {
		for {
			event := termbox.PollEvent()
			eventChan <- event
		}
	}()

	// register signals to channel
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// handle termbox events and unix signals
	func() { //++ TODO: dont use function literal. use labels instead.
		for {
			// select for either event or signal
			select {
			case event := <-eventChan:
				log.Printf("Have event: \n%s", spew.Sdump(event))
				// switch on event type
				switch event.Type {
				case termbox.EventKey: // actions depend on key
					switch event.Key {
					case termbox.KeyCtrlZ, termbox.KeyCtrlC:
						return
					}

					switch event.Ch {
					case 'q':
						return

					case 'w':
						scene.character.Turn(SNAKE_DIRECTION_UP)
					case 's':
						scene.character.Turn(SNAKE_DIRECTION_DOWN)
					case 'a':
						scene.character.Turn(SNAKE_DIRECTION_LEFT)
					case 'd':
						scene.character.Turn(SNAKE_DIRECTION_RIGHT)
					}

				case termbox.EventResize: // set sizes
					//setSizes(event.Width, event.Height)
					log.Println("size changed")

				case termbox.EventError: // quit
					log.Fatalf("Quitting because of termbox error: \n%s\n", event.Err)
				}
			case signal := <-sigChan:
				log.Printf("Have signal: \n%s", signal)
				return
			}
		}
	}()

	// close up
	termbox.Close()
	log.Println("stopping gosnake")
	os.Exit(0)
}
