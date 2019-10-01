package joe

import (
	"context"
	"fmt"

	"github.com/go-joe/joe/reactions"
)

// A Message is automatically created from a ReceiveMessageEvent and then passed
// to the RespondFunc that was registered via Bot.Respond(…) or Bot.RespondRegex(…)
// when the message matches the regular expression of the handler.
type Message struct {
	Context  context.Context
	Text     string
	AuthorID string
	Channel  string
	Matches  []string    // contains all sub matches of the regular expression that matched the Text
	Data     interface{} // corresponds to the ReceiveMessageEvent.Data field

	adapter Adapter
}

// Respond is a helper function to directly send a response back to the channel
// the message originated from. This function ignores any error when sending the
// response. If you want to handle the error use Message.RespondE instead.
func (msg *Message) Respond(text string, args ...interface{}) {
	_ = msg.RespondE(text, args...)
}

// RespondE is a helper function to directly send a response back to the channel
// the message originated from. If there was an error it will be returned from
// this function.
func (msg *Message) RespondE(text string, args ...interface{}) error {
	if len(args) > 0 {
		text = fmt.Sprintf(text, args...)
	}

	return msg.adapter.Send(text, msg.Channel)
}

func (msg *Message) React(reaction reactions.Reaction) (ok bool, err error) {
	adapter, ok := msg.adapter.(ReactionAwareAdapter)
	if !ok {
		return false, nil
	}

	return true, adapter.React(reaction, *msg)
}
