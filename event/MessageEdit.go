package event

import (
	"github.com/AlexBrin/goVkBot/vk/object"
)

type MessageEdit struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageEdit) GetName() string {
	return MessageEditEvent
}
