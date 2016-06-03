package curses

// Handles all output.

import (
	"github.com/discoviking/roguemike/io"
	"github.com/rthornton128/goncurses"
)

var screen *goncurses.Window
var Input chan *io.UpdateBundle

func Init() error {
	s, err := goncurses.Init()
	screen = s
	if err != nil {
		return err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)

	go func() {
		for s := range Input {
			output(s)
		}
	}()

	return nil
}

func Term() {
	goncurses.End()
}

func output(u *io.UpdateBundle) {
	clearscreen()
	for _, e := range u.Entities {
		draw(e)
	}
	refresh()
}

func clearscreen() {
	screen.Erase()
}

func refresh() {
	screen.Refresh()
}

func draw(e *io.EntityData) {
	screen.MoveAddChar(e.Y, e.X, 'X')
}
