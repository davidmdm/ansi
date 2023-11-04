package ansi

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

type Terminal struct {
	io.Writer
}

func (term Terminal) Print(args ...any) {
	_, _ = fmt.Fprint(term.Writer, args...)
}

func (term Terminal) Println(args ...any) {
	_, _ = fmt.Fprintln(term.Writer, args...)
}

func (term Terminal) Printf(format string, args ...any) {
	_, _ = fmt.Fprintf(term.Writer, format, args...)
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

func (term Terminal) Set(modes ...int) {
	fmt.Fprint(term.Writer, Esc(modes...))
}

type SpinnerOptions struct {
	Chars           []string
	InitialText     string
	SpinnerInterval time.Duration
	ClearAfterStop  bool
}

func (term Terminal) Spinner(opts SpinnerOptions) (chan<- string, func()) {
	term.SavePosition()

	if len(opts.Chars) == 0 {
		opts.Chars = []string{"|", "/", "-", "\\"}
	}

	if opts.SpinnerInterval == 0 {
		opts.SpinnerInterval = 650 * time.Millisecond
	}

	var (
		i        int
		wg       sync.WaitGroup
		text     = opts.InitialText
		charLen  = len(opts.Chars)
		messages = make(chan string)
		done     = make(chan struct{})
	)

	print := func() {
		term.ResetCursor()
		term.ClearAfterCursor()
		term.Printf("%s %s", opts.Chars[i%charLen], text)
	}

	print()

	wg.Add(1)
	go func() {
		ticker := time.NewTicker(opts.SpinnerInterval)
		defer ticker.Stop()

		for {
			select {
			case text = <-messages:
				print()
			case <-ticker.C:
				i++
				ticker.Reset(opts.SpinnerInterval)
				print()
			case <-done:
				wg.Done()
				return
			}
		}
	}()

	stop := func() {
		close(done)
		wg.Wait()
		if opts.ClearAfterStop {
			term.ResetCursor()
			term.ClearAfterCursor()
		}
	}

	return messages, stop
}

var (
	Stdout = Terminal{os.Stdout}
	Stderr = Terminal{os.Stderr}
)

func Esc(modes ...int) string {
	modeTxt := ""
	for i, mode := range modes {
		if i != 0 {
			modeTxt += ";"
		}
		modeTxt += strconv.Itoa(mode)
	}
	return "\u001B[" + modeTxt + "m"
}

const (
	ResetModes         = 0
	Bold               = 1
	Dim                = 2
	Italic             = 3
	Underline          = 4
	Blink              = 5
	Inverse            = 7
	Hidden             = 8
	StrikeThrough      = 9
	ResetBoldDim       = 22
	ResetItalic        = 23
	ResetUnderline     = 24
	ResetBlink         = 25
	ResetInverse       = 27
	ResetHidden        = 28
	ResetStrikethrough = 29
	FgBlack            = 30
	BgBlack            = 40
	FgRed              = 31
	BgRed              = 41
	FgGreen            = 32
	BgGreen            = 42
	FgYellow           = 33
	BgYellow           = 43
	FgBlue             = 34
	BgBlue             = 44
	FgMagenta          = 35
	BgMagenta          = 45
	FgCyan             = 36
	BgCyan             = 46
	FgWhite            = 37
	BgWhite            = 47
	FgDefault          = 39
	BgDefault          = 49
)
