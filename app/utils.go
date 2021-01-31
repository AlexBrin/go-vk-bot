package app

import (
	"context"

	"github.com/AlexBrin/go-vkbot/utils"
)

func ExtractBotTx(ctx context.Context) *Bot {
	return ctx.Value(utils.ContextKeyBot).(*Bot)
}
