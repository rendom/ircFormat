package ircFormat

import (
	"fmt"
	"testing"
)

func TestFgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x038%s\x03", text)

	c := New(text).SetFg(yellow).String()

	if c != expect {
		t.Fatalf("Got %s, expected %s", c, expect)
	}
}

func TestBgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x03,6%s\x03", text)

	c := New(text).SetBg(purple).String()

	if c != expect {
		t.Fatalf("Got %s, expected %s", c, expect)
	}
}

func TestBgAndFgColor(t *testing.T) {
	text := "test"
	expect := fmt.Sprintf("\x038,6%s\x03", text)

	c := New(text).SetFg(yellow).SetBg(purple).String()

	if c != expect {
		t.Fatalf("Got %s, expected %s", c, expect)
	}
}
