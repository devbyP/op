package greetingdomain

type GreetingMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Type    uint8  `json:"type"`
}
