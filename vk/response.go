package vk

import (
	"github.com/AlexBrin/goVkBot/vk/object"
)

type Response struct {
	Response interface{}
	Error    *object.ResponseError
	Raw      map[string]interface{}
}
