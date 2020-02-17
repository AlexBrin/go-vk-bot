package keyboard

import (
	"encoding/json"
	"github.com/AlexBrin/goVkBot/vk"
	"log"
	"strings"
	"text/template"
)

type Action struct {
	Type    string
	Payload string `json:",omitempty"`
	Link    string `json:",omitempty"`
	Label   string
}

type Button struct {
	Action Action
	Color  string `json:",omitempty"`
}

func New(inline bool, args ...[]Button) vk.H {
	h := vk.H{}
	params := vk.H{}

	h["inline"] = inline
	h["one_time"] = false
	h["buttons"] = args
	jsonKeyboard, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}
	params["keyboard"] = strings.ToLower(string(jsonKeyboard))
	return params
}

func ButtonText(label, command, color string, payload map[string]interface{}) Button {
	payload["command"] = command
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	b := Button{
		Action: Action{
			Type:    "text",
			Payload: string(jsonPayload),
			Label:   template.URLQueryEscaper(label),
		},
		Color: color,
	}
	return b
}

func ButtonLink(label, link string) Button {
	b := Button{
		Action: Action{
			Type:  "open_link",
			Link:  link,
			Label: template.URLQueryEscaper(label),
		},
	}

	return b
}
