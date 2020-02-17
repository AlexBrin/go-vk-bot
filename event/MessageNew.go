package event

import (
	"github.com/AlexBrin/goVkBot/vk/object"
)

type MessageNew struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageNew) GetName() string {
	return MessageNewEvent
}
