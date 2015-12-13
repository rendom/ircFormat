package ircFormat

import (
	"fmt"
	"testing"
)

func TestFgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x038%s\x03", text)

	c := New(text).SetFg(Yellow).String()

	if c != expect {
		t.Fatalf("Got %q, expected %q", c, expect)
	}
}

func TestBgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x03,6%s\x03", text)

	c := New(text).SetBg(Purple).String()

	if c != expect {
		t.Fatalf("Got %q, expected %q", c, expect)
	}
}

func TestBgAndFgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x038,6%s\x03", text)

	c := New(text).SetFg(Yellow).SetBg(Purple).String()

	if c != expect {
		t.Fatalf("Got %q, expected %q", c, expect)
	}
}

func TestChainAllMethods(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x038,6\x1D\x02\x1F%s\x1F\x02\x1D\x03", text)
	c := New(text).SetFg(Yellow).SetBg(Purple).SetBold().SetItalic().SetUnderline().String()

	if c != expect {
		t.Fatalf("Got %q, expected %q", c, expect)
	}
}

func TestUnderlineText(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x1F%s\x1F", text)

	c := UnderlineText(text)
	if c != expect {
		t.Fatalf("Got %q, expected %q", c, expect)
	}
}

func TestBoldText(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x02%s\x02", text)

	c := BoldText(text)
	if c != expect {
		t.Fatalf("Got %q, expected %q", c, expect)
	}
}

func TestItalicText(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x1D%s\x1D", text)

	c := ItalicText(text)
	if c != expect {
		t.Fatalf("Got %q, expected %q", c, expect)
	}
}
