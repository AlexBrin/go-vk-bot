package main

import (
	"context"
	"fmt"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"

	internalEvents "github.com/AlexBrin/go-vkbot/events"

	"github.com/AlexBrin/go-vkbot/app"

	"github.com/SevereCloud/vksdk/v2/events"
)

func main() {
	bot := app.New()

	bot.GetFuncList().CommandNew("test", func(ctx context.Context, command internalEvents.CommandNew) {
		bot := app.ExtractBotTx(ctx)

		_, _ = bot.SendMessage(command.Object.Message.PeerID, "command handler", nil)

		fmt.Printf("%#v\n", command)
	})

	bot.GetFuncList().CommandNew("keyboard", func(ctx context.Context, command internalEvents.CommandNew) {

		keyboard := object.NewMessagesKeyboard(false)
		keyboard.
			AddRow().
			AddTextButton("Test", "inline-keyboard", "primary").
			AddRow().
			AddCallbackButton("second", "payload", "positive")

		_, _ = bot.SendMessage(command.Object.Message.PeerID, "keyboard for you", &api.Params{
			"keyboard": keyboard.ToJSON(),
		})
	})

	bot.GetFuncList().CommandNew("inline-keyboard", func(ctx context.Context, command internalEvents.CommandNew) {
		keyboard := object.NewMessagesKeyboardInline()
		keyboard.
			AddRow().
			AddTextButton("Test", "!inline-keyboard", "primary").
			AddRow().
			AddLocationButton("!location")

		_, _ = bot.SendMessage(command.Object.Message.PeerID, "inline keyboard for you", &api.Params{
			"keyboard": keyboard.ToJSON(),
		})
	})

	bot.GetFuncList().MessageNew(func(ctx context.Context, object events.MessageNewObject) {
		bot := app.ExtractBotTx(ctx)

		if bot.MessageIsCommand(object) {
			return
		}

		fmt.Printf("%#v\n", *bot)
		fmt.Println(object.Message.Text)

		_, _ = bot.SendMessage(object.Message.FromID, "Привет", nil)
	})

	bot.Polling()
}
