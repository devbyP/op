package model

type (
	MessageType uint8

	MessageData struct {
		ID        int         `json:"id"`
		Message   string      `json:"message"`
		Type      MessageType `json:"type"`
		CreatedAt int64       `json:"created_at"`
	}
)

// message type enum
const (
	Normal MessageType = iota
	Casual
	Polite
)

var (
	messageTypeString = []string{"Normal", "Casual", "Polite"}
)

func (m MessageType) String() string {
	if int(m) > len(messageTypeString) {
		return "Unknown"
	}
	return messageTypeString[m]
}
