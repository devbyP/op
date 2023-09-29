package greetingapi

import (
	"op/greetingdomain"
)

type GreetingResponse struct {
	ID   int                             `json:"id"`
	Data *greetingdomain.GreetingMessage `json:"data"`
}

type GreetingAllResponse struct {
	Data []*greetingdomain.GreetingMessage `json:"data"`
}

type SaveNewGreetingRequest struct {
	Message string `json:"message"`
	Type    int    `json:"type"`
}

type SaveNewGreetingResponse struct {
	ID int `json:"id"`
}
