package app

import (
	"context"

	internalEvents "github.com/AlexBrin/go-vkbot/events"

	"github.com/SevereCloud/vksdk/v2/events"
)

type CommandHandler func(ctx context.Context, command internalEvents.CommandNew)

type FuncList struct {
	events.FuncList
	goroutine bool

	commandHandlers map[string][]CommandHandler
}

func (fl *FuncList) Goroutine(v bool) {
	fl.FuncList.Goroutine(v)
	fl.goroutine = v
}

func (fl *FuncList) CommandNew(command string, handler CommandHandler) {
	if _, ok := fl.commandHandlers[command]; !ok {
		fl.commandHandlers[command] = make([]CommandHandler, 0, 1)
	}

	fl.commandHandlers[command] = append(fl.commandHandlers[command], handler)
}

func (fl *FuncList) HandleCommand(ctx context.Context, commandNew internalEvents.CommandNew) {
	if _, ok := fl.commandHandlers[commandNew.Command]; !ok {
		return
	}

	for _, f := range fl.commandHandlers[commandNew.Command] {
		if fl.goroutine {
			go f(ctx, commandNew)
		} else {
			f(ctx, commandNew)
		}
	}
}
