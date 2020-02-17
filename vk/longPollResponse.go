package vk

type LongPollUpdate struct {
	EventType string      `json:"type"`
	Object    interface{} `json:"object"`
	GroupId   float64     `json:"group_id"`
}

type LongPollResponse struct {
	TS      string           `json:"ts"`
	Updates []LongPollUpdate `json:"updates"`
	Failed  float64          `json:"failed"`
}
