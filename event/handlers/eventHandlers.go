package handlers

import "github.com/AlexBrin/goVkBot/event"

type EventHandler func(event.Event) bool

// messages
type CommandHandler func(args []string, command *event.Command) bool
type PayloadHandler func(payload map[string]string, command *event.Payload) bool
type MessageNewHandler func(messageNew *event.MessageNew) bool
type MessageEditHandler func(edit *event.MessageEdit) bool
type MessageReplyHandler func(reply *event.MessageReply) bool
type MessageAllowHandler func(allow *event.MessageAllow) bool
type MessageDenyHandler func(deny *event.MessageDeny) bool
