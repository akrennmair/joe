package main

import (
	"github.com/go-joe/joe"
	"github.com/go-joe/joe/examples/react/reactions"
)

type ExampleBot struct {
	*joe.Bot
	Reactions *reactions.Module
}

func main() {
	b := NewBot("xoxb-1452345…")
	err := b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func NewBot(slackToken string) *ExampleBot {
	b := &ExampleBot{
		Reactions: reactions.New(),
	}

	b.Bot = joe.New("example",
		// slack.Adapter(slackToken),
		b.Reactions.Module(),
	)

	b.Respond("hello", b.MyHandler)
	b.RegisterHandler(b.ReceiveReaction)

	return b
}

func (b *ExampleBot) MyHandler(msg joe.Message) error {
	msg.Respond("OK")
	return b.Reactions.Add(reactions.ThumbsUp, msg)
}

func (b *ExampleBot) ReceiveReaction(evt reactions.Event) error {
	return nil
}
