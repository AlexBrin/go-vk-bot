package app

import (
	"context"
	"os"
	"os/signal"
	"strings"

	internalEvents "github.com/AlexBrin/go-vkbot/events"

	"github.com/AlexBrin/go-vkbot/utils"

	"github.com/SevereCloud/vksdk/v2/events"

	"github.com/AlexBrin/go-vkbot/config"
	"github.com/AlexBrin/go-vkbot/logger"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/spf13/pflag"
)

// Bot struct.
type Bot struct {
	client *api.VK
	conf   *config.Configuration

	commandPrefixes []string

	FuncList
}

// New Возвращает новый экземпляр Bot.
func New() *Bot {
	filePath := pflag.String("config-path", "", "Path to config file")
	vkToken := pflag.String("vk-token", "", "VK API Token")
	pflag.Parse()

	conf, err := config.NewConf(*filePath)
	if err != nil {
		logger.Message{Message: "Ошибка инициализации конфигурации", Err: err}.Panic()
	}

	conf.Conf.SetDefault(config.VkToken, vkToken)

	funcList := FuncList{
		FuncList:        *events.NewFuncList(),
		commandHandlers: map[string][]CommandHandler{},
	}

	funcList.MessageNew(func(ctx context.Context, object events.MessageNewObject) {
		bot := ExtractBotTx(ctx)
		message := object.Message.Text

		if bot.IsCommand(message) {
			args := strings.Split(message[1:], " ")
			funcList.HandleCommand(ctx, internalEvents.CommandNew{
				Command:   args[0],
				Arguments: args[1:],
				Object:    object,
			})
		}
	})

	funcList.MessageNew(func(ctx context.Context, object events.MessageNewObject) {
		bot := ExtractBotTx(ctx)

		payload := object.Message.Payload
		if bot.IsCommand(payload) {
			args := strings.Split(payload[1:], " ")
			funcList.HandleCommand(ctx, internalEvents.CommandNew{
				Command:   args[0],
				Arguments: args[1:],
				Object:    object,
			})
		}
	})

	return &Bot{
		client:   api.NewVK(conf.Conf.GetString(config.VkToken)),
		conf:     conf,
		FuncList: funcList,

		commandPrefixes: conf.Conf.GetStringSlice(config.VkCommandPrefix),
	}
}

// AddCommandPrefix Добавляет префикс команд в рантайме.
func (b *Bot) AddCommandPrefix(prefix string) *Bot {
	b.commandPrefixes = append(b.commandPrefixes, prefix)

	return b
}

// GetCommandPrefixes Возвращает префиксы команд.
func (b *Bot) GetCommandPrefixes() []string {
	return b.commandPrefixes
}

// GetClient Возвращает клиент ВК.
func (b *Bot) GetClient() *api.VK {
	return b.client
}

// GetFuncList Возвращает текущий список функций.
func (b *Bot) GetFuncList() *FuncList {
	return &b.FuncList
}

func (b *Bot) MessageIsCommand(message events.MessageNewObject) bool {
	return b.IsCommand(message.Message.Text) || b.IsCommand(message.Message.Payload)
}

// IsCommand Проверяет, является ли переданное сообщение командой.
func (b *Bot) IsCommand(message string) bool {
	if len(message) > 1 {
		prefix := message[:1]
		for _, availablePrefix := range b.commandPrefixes {
			if prefix == availablePrefix {
				return true
			}
		}
	}

	return false
}

// SendMessage Отправляет сообщение в указанный чат.
func (b *Bot) SendMessage(chatId int, message string, params *api.Params) (messageId int, err error) {
	if params == nil {
		params = &api.Params{}
	}

	(*params)["message"] = message
	(*params)["peer_id"] = chatId
	(*params)["random_id"] = utils.GetRandomMessageId()

	messageId, err = b.client.MessagesSend(*params)

	if err != nil {
		logger.Message{
			Message: "Ошибка при отправке сообщения",
			Err:     err,
		}.AddField("peer_id", chatId).Error()
	}

	return messageId, err
}

// Polling Запускает прослушивание LongPoll сервера.
func (b *Bot) Polling() {
	logger.Message{Message: "Запуск LongPoll..."}.Info()

	lp, err := longpoll.NewLongPoll(b.client, b.conf.Conf.GetInt(config.VkGroupId))
	if err != nil {
		logger.Message{Message: "Ошибка при получении LongPoll сервера", Err: err}.Panic()
	}

	lp.FuncList = b.FuncList.FuncList

	lp.Goroutine(b.conf.Conf.GetBool(config.BotGoroutine))

	ctx := context.Background()
	ctx = context.WithValue(ctx, utils.ContextKeyBot, b)

	var done bool
	go func() {
		err = lp.RunWithContext(ctx)
		done = true
	}()

	logger.Message{Message: "Работаю"}.Info()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

ctxWatcher:
	for {
		if done {
			lp.Shutdown()
			logger.Message{Message: "LongPoll завершил работу"}.Info()
			if err != nil && err.Error() != "context canceled" {
				logger.Message{Message: "Произошла ошибка при выполнении", Err: err}.Panic()
			}
		}

		select {
		case <-ch:
			logger.Message{Message: "Завершение работы"}.Info()
			lp.Shutdown()
			break ctxWatcher
		default:
			continue
		}
	}
}
