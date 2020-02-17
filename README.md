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
    "github.com/AlexBrin/goVkBot/keyboard"

)

const (
	groupID = "group id"
	token   = "group token"
	version = "5.103"
)

var bot *govkbot.Bot

func main() {
	bot = govkbot.CreateBot(groupID, token, version)

	// Handling `test` command
	bot.OnCommand("test", func(args []string, command *event.Command) bool {

        //example  keyboard
        g := keyboard.ButtonText("🕹 Play", "game", "positive", vk.H{"count": "100"})
		s := keyboard.ButtonLink("Url", "https://vk.com/im")
		params := keyboard.New(true, []keyboard.Button{g, g, g, g}, []keyboard.Button{s}, []keyboard.Button{g, s})

		bot.SendMessage("Hello!", command.PrivateMessage.Message.PeerID, params)
		return true
	})
    
    // Handling `game` payload
    bot.OnPayload("game", func(payload map[string]string, ev *event.Payload) bool {
		bot.SendMessage(payload["count"], ev.PrivateMessage.Message.UserID, vk.H{})
		return true
	})

	// Handling of all messages except commands and payloads
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
