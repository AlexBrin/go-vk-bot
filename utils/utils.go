package utils

import (
	"time"
)

// GetRandomMessageId возвращает random_id для сообщений.
func GetRandomMessageId() int32 {
	return int32(time.Now().UnixNano())
}
