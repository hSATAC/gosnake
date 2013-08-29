package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
    "os"
    "os/signal"
    "log"
    "time"
    "math/rand"
)

func main() {
    // initialize termbox
    err := termbox.Init()
    if err != nil {
        fmt.Println("Could not start termbox for gomatrix. View ~/.gomatrix-log for error messages.")
        log.Printf("Cannot start gomatrix, termbox.Init() gave an error:\n%s\n", err)
        os.Exit(1)
    }
    termbox.HideCursor()

    // go
    go func() {
        for {
            <-time.After(40 * time.Millisecond) //++ TODO: find out wether .After() or .Sleep() is better performance-wise
            termbox.Flush()
            //fmt.Printf("█")
            x := rand.Intn(8)
            termbox.SetCell(x*2-1, rand.Intn(8),' ', termbox.ColorDefault, termbox.ColorRed)
            termbox.SetCell(x*2-2, rand.Intn(8),' ', termbox.ColorDefault, termbox.ColorRed)

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
                //log.Printf("Have event: \n%s", spew.Sdump(event))
                // switch on event type
                switch event.Type {
                case termbox.EventKey: // actions depend on key
                    switch event.Key {
                    case termbox.KeyCtrlZ, termbox.KeyCtrlC:
                        return
                        //++ TODO: add more fun keys (slowmo? freeze? rampage?)
                    }

                    switch event.Ch {
                    case 'q':
                        return

                    case 'c':
                        termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

                    case 'a':
                        //characters = alphaNumerics

                    case 'k':
                        //characters = halfWidthKana
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
    log.Println("stopping gomatrix")
    fmt.Println("Thank you for connecting with Morpheus' Matrix API v4.2. Have a nice day!")
    os.Exit(0)
}

