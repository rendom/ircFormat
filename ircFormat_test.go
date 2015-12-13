package ircFormat

import (
	"fmt"
	"testing"
)

func TestFgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x038%s\x03", text)

	c := New(text).setFg(yellow).String()

	if c != expect {
		t.Fatalf("Got %s, expected %s", c, expect)
	}
}

func TestBgColor(t *testing.T) {
	expect := fmt.Sprintf("\x03,6%s\x03", text)
	text := "test"

	c := New(text).setBg(purple).String()

	if c != expect {
		t.Fatalf("Got %s, expected %s", c, expect)
	}
}

func TestBgAndFgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x038,6%s\x03", text)

	c := New(text).setFg(yellow).setBg(purple).String()

	if c != expect {
		t.Fatalf("Got %s, expected %s", c, expect)
	}
}
