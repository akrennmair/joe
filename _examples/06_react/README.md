# React Feature

## Use cases
- add reaction to messages
- trigger handler when reaction is added
- list reactions on a message

# How to add reactions to different adapters
- slack: `client.AddReactionContext(ctx, reaction, slack.NewRefToMessage(channelID, msgTimestamp))`
    - requires a message timestamp
    - the timestamp of a message is the unique per-channel and serves as the messages identifier within that channel
- rocket chat: `ReactToMessage(message *models.Message, reaction string) error`
    - requires a message ID
- telegram: unclear but apparently possible (editing messages HTML?)
    - probably too complicated to actually be implemented by the Adapter (not normally a thing on telegram?)
- signal: none
    - no concept of reactions
- twitter: none
    - no concept of reactions, only Liking (heart)

# How did other bot libraries implement reactions?
- feature set?
- where did they struggle?

* https://github.com/sgreben/telegram-emoji-reactions-bot

# Would the chosen approah also work nicely with other adapter specific features?
