package events

import "github.com/SevereCloud/vksdk/v2/events"

type CommandNew struct {
	Command   string
	Arguments []string
	Object    events.MessageNewObject
}
