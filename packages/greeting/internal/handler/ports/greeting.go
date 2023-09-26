package ports

import (
	"context"

	"op/greetingapi"
)

type IGreeting interface {
	Greeting(ctx context.Context) (*greetingapi.GreetingResponse, error)
	SaveMessage(ctx context.Context, req *greetingapi.SaveNewGreetingRequest) error
}
