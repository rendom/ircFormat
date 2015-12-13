package ircFormat

import "strconv"

const (
	White = iota
	Black
	Blue
	Green
	Red
	Brown
	Purple
	Orange
	Yellow
	Lime
	Teal
	Cyan
	Darkblue
	Pink
	Grey
	Lightgrey
	None
)

const (
	resetFormatChar = "\x0F"
	boldChar        = "\x02"
	italicChar      = "\x1D"
	underlineChar   = "\x1F"
	swapChar        = "\x16"
	colorChar       = "\x03"
)

type IrcText struct {
	text      string
	bgColor   int
	fgColor   int
	bold      bool
	italic    bool
	underline bool
	swap      bool
}

// New creates IrcText object and take your text as argument.
func New(s string) *IrcText {
	return &IrcText{text: s, bgColor: None, fgColor: None}
}

// SetFg sets foreground color
func (i *IrcText) SetFg(c int) *IrcText {
	i.fgColor = c
	return i
}

// SetBg sets background color
func (i *IrcText) SetBg(c int) *IrcText {
	i.bgColor = c
	return i
}

func (i *IrcText) String() string {
	if i.bgColor == None && i.fgColor == None {
		return i.text
	}

	c := colorChar

	if i.fgColor != None {
		c += strconv.Itoa(i.fgColor)
	}

	if i.bgColor != None {
		c += ","
		c += strconv.Itoa(i.bgColor)
	}

	c += i.text
	c += colorChar

	return c
}
