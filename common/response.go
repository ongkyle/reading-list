package common

import (
	"time"
)

type Response struct {
	ID        uint64 `json:"id,omitempty"`
	Data      []Item `json:"data,omitempty"`
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

func NewResponse(code int, data []Item) Response {
	response := Response{
		Code:      code,
		Data:      data,
		Timestamp: getCurrentUnixTime(),
	}
	return response
}

func getCurrentUnixTime() int64 {
	return int64(time.Now().Unix())
}