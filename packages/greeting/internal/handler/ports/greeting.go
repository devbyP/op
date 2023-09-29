package ports

import (
	"op/greetingapi"
)

type IGreeting interface {
	Greeting(id int) (*greetingapi.GreetingResponse, error)
	GreetingAll() (*greetingapi.GreetingAllResponse, error)
	SaveMessage(req *greetingapi.SaveNewGreetingRequest) (*greetingapi.SaveNewGreetingResponse, error)
}
