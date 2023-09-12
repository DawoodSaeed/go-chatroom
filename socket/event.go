package socket

import (
	"encoding/json"
)

const (
	EventMessageType = "message_sent"
)

// This is what we will recieve ;
type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// The function that will be executing according to the event;
type EventHandlerFn func(e *Event, client *Client) error

// This is what we will send to the client;
type EventSendMessage struct {
	Message string `json:"message"`
	From    string `json:"from"`
}
