package ansi

import (
	"fmt"
	"io"
	"os"
)

type Terminal struct {
	io.Writer
}

func (term Terminal) SavePosition() {
	_, _ = fmt.Fprint(term.Writer, "\u001B7")
}

func (term Terminal) ResetCursor() {
	_, _ = fmt.Fprint(term.Writer, "\u001B8")
}

func (term Terminal) ClearAfterCursor() {
	_, _ = fmt.Fprint(term.Writer, "\u001B[0J")
}

func (term Terminal) ClearBeforeCursor() {
	_, _ = fmt.Fprint(term.Writer, "\u001B[1J")
}

func (term Terminal) ClearScreen() {
	_, _ = fmt.Fprint(term.Writer, "\u001B[2J")
}

func (term Terminal) ClearLineAfterCursor() {
	_, _ = fmt.Fprint(term.Writer, "\u001B[0K")
}

func (term Terminal) ClearLineBeforeCursor() {
	_, _ = fmt.Fprint(term.Writer, "\u001B[1K")
}

func (term Terminal) ClearLine() {
	_, _ = fmt.Fprint(term.Writer, "\u001B[2K")
}

var (
	Stdout = Terminal{os.Stdout}
	Stderr = Terminal{os.Stderr}
)
