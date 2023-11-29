package ansi

import (
	"fmt"
	"io"
	"strconv"
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

var resetModesSeq = Esc(ResetModes)

type Style struct {
	esc string
}

func MakeStyle(modes ...int) Style {
	return Style{
		esc: Esc(modes...),
	}
}

func (style Style) sprint(value string) string {
	return fmt.Sprintf("%s%s%s", style.esc, value, resetModesSeq)
}

func (style Style) Sprint(args ...any) string {
	return style.sprint(fmt.Sprint(args...))
}

func (style Style) Sprintln(args ...any) string {
	return style.sprint(fmt.Sprintln(args...))
}

func (style Style) Sprintf(format string, args ...any) string {
	return style.sprint(fmt.Sprintf(format, args...))
}

func (style Style) Fprintf(w io.Writer, format string, args ...any) {
	_, _ = fmt.Fprint(w, style.Sprintf(format, args...))
}

func (style Style) Fprint(w io.Writer, args ...any) {
	_, _ = fmt.Fprint(w, style.Sprint(args...))
}

func (style Style) Fprintln(w io.Writer, args ...any) {
	_, _ = fmt.Fprint(w, style.Sprintln(args...))
}

func (style Style) Printf(format string, args ...any) {
	_, _ = fmt.Print(style.Sprintf(format, args...))
}

func (style Style) Print(args ...any) {
	_, _ = fmt.Print(style.Sprint(args...))
}

func (style Style) Println(args ...any) {
	_, _ = fmt.Print(style.Sprintln(args...))
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
