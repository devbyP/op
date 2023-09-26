package greetingdomain

type GreetingMessage struct {
	Message string `json:"message"`
	Type    uint8  `json:"type"`
}
