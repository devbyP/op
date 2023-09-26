package greetingapi

import (
	"op/greetingdomain"
)

type GreetingResponse struct {
	ID   int                            `json:"id"`
	Data greetingdomain.GreetingMessage `json:"data"`
}

type SaveNewGreetingRequest struct {
	Data greetingdomain.GreetingMessage `json:"data"`
}
