package main

import (
	"log"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

var s tcell.Screen
var row = 0
var col = 0
var command = []rune{}

func EventLoop(quit chan<- struct{}) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				close(quit)
				return
			case tcell.KeyEnter:
				// TODO run the command
				NewLine()
				s.Sync()
			// TODO implement backspace
			case tcell.KeyDEL:
				DeleteRune()
				s.Sync()
			default:
				// TODO collect command input
				PrintRune(ev.Rune())
				s.Sync()
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}

func PrintRune(r rune) {
	// TODO handle line wraps
	width := runewidth.RuneWidth(r)
	s.SetContent(col, row, r, nil, tcell.StyleDefault)
	col += width
	command = append(command, r)
}

func DeleteRune() {
	// TODO handle deleting from line wraps
	if col == 0 {
		return
	}

	command = command[:len(command)-1]
	col--
	PrintRune(' ')
	col--
}

func NewLine() {
	row++
	col = 0
	command = nil
}

func main() {
	var err error
	s, err = tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}
	defer s.Fini()

	quit := make(chan struct{})
	go EventLoop(quit)

	<-quit
}
