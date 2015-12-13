[![Build Status](https://travis-ci.org/rendom/ircFormat.svg?branch=master)](https://travis-ci.org/rendom/ircFormat)

# ircFormat
Library to format irc messages

https://godoc.org/github.com/rendom/ircFormat


####Example usage:####
```
// Bold text with red foreground
i.Privmsg(CHAN, ircFormat.New("Test").SetBold().SetFg(ircFormat.Red));

// Bold text
i.Privmsg(CHAN, ircFormat.BoldText("Test!"));

// Italic text
i.Privmsg(CHAN, ircFormat.ItalicText("Test!"));

// Underline text
i.Privmsg(CHAN, ircFormat.UnderlineText("Test!"));

// Red text
i.Privmsg(CHAN, ircFormat.ColorizeText("Test!", ircFormat.Red, ircFormat.None));

// Red text with green background
i.Privmsg(CHAN, ircFormat.ColorizeText("Test!", ircFormat.Red, ircFormat.Green));
```
