package govkbot

import (
	"encoding/json"
	"fmt"
	"github.com/AlexBrin/goVkBot/event"
	"github.com/AlexBrin/goVkBot/event/handlers"
	"github.com/AlexBrin/goVkBot/log"
	"github.com/AlexBrin/goVkBot/vk"
	"github.com/AlexBrin/goVkBot/vk/object"
	"github.com/mitchellh/mapstructure"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Bot struct {
	groupId string
	token   string
	logger  *log.Log
	api     *vk.API

	prefixList []string

	commandHandlers map[string][]handlers.CommandHandler
	payloadHandlers map[string][]handlers.PayloadHandler
	handlers        map[string][]handlers.EventHandler
}

func createDecoder(output interface{}) (decoder *mapstructure.Decoder) {
	decoder, _ = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		TagName:  "map",
		Result:   output,
	})
	return
}

func CreateBot(groupId, token, version string) (b *Bot) {
	b = &Bot{
		groupId: groupId,
		token:   token,
		logger:  log.Create("[%s] %s"),
		api:     vk.Create(token, version),

		prefixList: []string{".", "/"},

		commandHandlers: map[string][]handlers.CommandHandler{},
		payloadHandlers: map[string][]handlers.PayloadHandler{},
		handlers:        map[string][]handlers.EventHandler{},
	}

	return
}

func (b *Bot) isPrefix(command string) bool {
	firstSymbol := command[:1]

	for _, prefix := range b.prefixList {
		if firstSymbol == prefix {
			return true
		}
	}

	return false
}

func (b *Bot) SetPrefixList(prefixes []string) {
	b.prefixList = prefixes
}

func (b *Bot) AddPrefix(prefix string) {
	if !b.isPrefix(prefix) {
		b.prefixList = append(b.prefixList, prefix)
	}
}

func (b *Bot) GetPrefixList() []string {
	return b.prefixList
}

func (b *Bot) GetLogger() *log.Log {
	return b.logger
}

func (b *Bot) GetApi() *vk.API {
	return b.api
}

func (b *Bot) handlersExists(eventType string) bool {
	_, ok := b.handlers[eventType]
	return ok
}

func (b *Bot) commandExists(command string) bool {
	_, ok := b.commandHandlers[strings.ToLower(command)]
	return ok
}

func (b *Bot) OnCommand(command string, h ...handlers.CommandHandler) {
	command = strings.ToLower(command)
	if !b.commandExists(command) {
		b.commandHandlers[command] = make([]handlers.CommandHandler, 0)
	}

	b.commandHandlers[command] = append(b.commandHandlers[command], h...)
}

func (b *Bot) payloadExists(payload string) bool {
	_, ok := b.payloadHandlers[strings.ToLower(payload)]
	return ok
}

func (b *Bot) OnPayload(payload string, h ...handlers.PayloadHandler) {
	payload = strings.ToLower(payload)
	if !b.payloadExists(payload) {
		b.payloadHandlers[payload] = make([]handlers.PayloadHandler, 0)
	}

	b.payloadHandlers[payload] = append(b.payloadHandlers[payload], h...)
}

func (b *Bot) On(eventType string, h ...handlers.EventHandler) {
	if !b.handlersExists(eventType) {
		b.handlers[eventType] = make([]handlers.EventHandler, 0)
	}

	b.handlers[eventType] = append(b.handlers[eventType], h...)
}

//func (b *Bot) executeEvent(eventType string, data interface{}) {
//	eventType = strings.ToLower(eventType)
//	switch eventType {
//
//	case event.CommandEvent:
//		ev := e.(event.Command)
//		if b.commandExists(ev.Command) {
//			for _, handler := range b.commandHandlers[ev.Command] {
//				handler(ev.Args, ev)
//			}
//		}
//
//	default:
//		b.handlersExists(eventType)
//		for _, handler := range b.handlers[eventType] {
//			handler(e.(event.Event))
//		}
//	}
//}

func (b *Bot) SendMessage(message string, to float64, params vk.H) {
	params["peer_id"] = to
	params["message"] = template.URLQueryEscaper(message)
	params["random_id"] = time.Now().UnixNano() + int64(to)

	response, err := b.api.Api("messages.send", params)
	if err != nil {
		b.GetLogger().Error(err.Error())
		return
	}

	if response.Error != nil {
		b.GetLogger().Error("Error:", strconv.Itoa(int(response.Error.ErrorCode)))
		b.GetLogger().Error(response.Error.Error)
	}
}

func (b *Bot) handle(updates []vk.LongPollUpdate) {
	var ev event.Event
	for _, update := range updates {
		switch update.EventType {

		case event.MessageNewEvent:
			pm := object.PrivateMessage{}
			_ = createDecoder(&pm).Decode(update.Object)

			args := strings.Split(pm.Message.Text, " ")
			if len(args) >= 1 && args[0] != "" {
				cmd := strings.ToLower(args[0])
				if b.isPrefix(cmd) {
					cmd = cmd[1:]
				}

				var next bool

				if b.commandExists("*") {
					for _, handler := range b.commandHandlers["*"] {
						if b.commandExists(cmd) {
							next = handler(args[1:], &event.Command{Command: args[0], Args: args[1:], PrivateMessage: &pm})
							if !next {
								break
							}
						}
					}
				}

				if pm.Message.Payload != "" {
					byt := []byte(pm.Message.Payload)
					var payload map[string]string
					if err := json.Unmarshal(byt, &payload); err != nil {
						b.GetLogger().Error("Error:", "Payload not json format!")
						return
					}
					if b.payloadExists("*") {
						for _, handler := range b.payloadHandlers["*"] {
							next = handler(payload, &event.Payload{Payload: payload, PrivateMessage: &pm})
							if !next {
								break
							}
						}
					}

					if b.payloadExists(payload["command"]) {
						for _, handler := range b.payloadHandlers[payload["command"]] {
							next = handler(payload, &event.Payload{Payload: payload, PrivateMessage: &pm})
							if !next {
								break
							}
						}

						continue
					}
				}

				if b.commandExists(cmd) {
					for _, handler := range b.commandHandlers[cmd] {
						next = handler(args[1:], &event.Command{Command: args[0], Args: args[1:], PrivateMessage: &pm})
						if !next {
							break
						}
					}

					continue
				}
			}

			ev = &event.MessageNew{PrivateMessage: &pm}

		case event.MessageEditEvent:
			pm := object.PrivateMessage{}
			_ = createDecoder(&pm).Decode(update.Object)
			ev = &event.MessageEdit{PrivateMessage: &pm}

		case event.MessageReplyEvent:
			pm := object.PrivateMessage{}
			_ = createDecoder(&pm).Decode(update.Object)
			ev = &event.MessageReply{PrivateMessage: &pm}

		case event.MessageAllowEvent:
			ev = &event.MessageAllow{}

		case event.MessageDenyEvent:
			ev = &event.MessageDeny{}

		}

		if ev == nil {
			continue
		}

		_ = createDecoder(ev).Decode(update.Object)
		var next bool
		for _, handle := range b.handlers[ev.GetName()] {
			next = handle(ev)
			if !next {
				break
			}
		}
	}
}

func (b *Bot) Polling() {
	b.GetLogger().Log("Getting LongPoll Server...")

	resp, err := b.GetApi().Api("groups.getLongPollServer", vk.H{
		"group_id": b.groupId,
	})
	if err != nil {
		panic(err)
	}

	var key, server, ts string
	respMap := resp.Response.(map[string]interface{})
	key = respMap["key"].(string)
	server = respMap["server"].(string)
	ts = respMap["ts"].(string)

	b.GetLogger().Info("Server:", server)

	for {
		lpResponse, err := http.Get(fmt.Sprintf("%s?act=a_check&key=%s&wait=1&ts=%s", server, key, ts))
		if err != nil {
			b.logger.Error(err.Error())
			b.logger.Error("Waiting 5 seconds...")
			time.Sleep(time.Second * 5)
			continue
		}

		bytes, err := ioutil.ReadAll(lpResponse.Body)
		if err != nil {
			continue
		}
		err = lpResponse.Body.Close()
		if err != nil {
			continue
		}

		response := vk.LongPollResponse{}
		err = json.Unmarshal(bytes, &response)
		if err != nil {
			b.logger.Error(err.Error())
			continue
		}

		switch {

		case response.Failed == 1:
			ts = response.TS
			continue

		case response.Failed == 2 || response.Failed == 3:
			resp, err := b.GetApi().Api("groups.getLongPollServer", vk.H{
				"group_id": b.groupId,
			})
			if err != nil {
				panic(err)
			}
			respMap := resp.Response.(map[string]interface{})
			key = respMap["key"].(string)
			server = respMap["server"].(string)
			ts = respMap["ts"].(string)

			continue

		}

		ts = response.TS

		if len(response.Updates) == 0 {
			continue
		}

		b.handle(response.Updates)
	}
}
