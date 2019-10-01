package main

import (
	"fmt"

	"github.com/go-joe/joe"
	"github.com/go-joe/joe/reactions"
)

func main() {
	b := joe.New("example-bot")
	b.Respond("hello", MyHandler)
	b.Brain.RegisterHandler(ReceiveReaction)

	err := b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func MyHandler(msg joe.Message) error {
	ok, err := msg.React(reactions.ThumbsUp)
	if !ok {
		msg.Respond("Sorry but I do not support reactions")
	}

	return err
}

func ReceiveReaction(evt reactions.Event) error {
	fmt.Printf("Received event: %+v", evt)
	return nil
}
