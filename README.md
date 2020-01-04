GoVkBot
=======

Module for fast development of bots for [VK groups](https://vk.com/dev/manuals) based on long polls

Example:
```go
package main

import (
	"github.com/AlexBrin/goVkBot"
	"github.com/AlexBrin/goVkBot/event"
	"github.com/AlexBrin/goVkBot/vk"
)

const (
	groupID = "group id"
	token   = "group token"
	version = "5.103"
	prefix  = "MyAwesomeBot"
)

var bot *govkbot.Bot

func main() {
	bot = govkbot.CreateBot(groupID, token, version)

	// Handling `test` command
	bot.OnCommand("test", func(args []string, command *event.Command) bool {
		bot.SendMessage("Hello!", command.PrivateMessage.Message.PeerID, vk.H{})
		return true
	})

	// Handling of all messages except commands
	bot.On(event.MessageNewEvent, func(e event.Event) bool {
		// This is necessary to get rid of a heap of methods like `OnNewWallReply` or `OnEditWallReply`
		ev := e.(*event.MessageNew)

		bot.GetLogger().Info(ev.GetName())
		return true
	})

	// Starting the bot and waiting the events
	bot.Polling()
}
```

Available events (for the method `On`):
```
MessageNewEvent - input new message
MessageReplyEvent - output message"
MessageEditEvent - edit message
MessageAllowEvent - subscribe on messages
MessageDenyEvent - unsubscribe from messages
```
