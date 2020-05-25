package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

const animationSpeed = 10 * time.Millisecond

func main()  {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	render()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		default:
			render()
			time.Sleep(animationSpeed)
		}
	}
}

const backgroundColor = termbox.ColorBlue
const instructionsColor = termbox.ColorYellow

func render() {
	termbox.Clear(backgroundColor, backgroundColor)
	tbprint(1, 1, instructionsColor, backgroundColor, "Test")
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}