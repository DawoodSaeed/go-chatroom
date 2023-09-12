package socket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type ClientMap map[*Client]bool

type Client struct {
	conn        *websocket.Conn
	socket      *Socket
	messagesChn chan []byte
}

// Returns a client;
func NewClient(connection *websocket.Conn, socket *Socket) *Client {
	return &Client{
		conn:        connection,
		socket:      socket,
		messagesChn: make(chan []byte),
	}
}

// For reading the messages;
func (client *Client) ReadMessages() {
	defer func() {
		client.socket.RemoveClient(client)
	}()

	for {
		_, payload, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Fatal("There is an unexpected error;")
			}
			log.Fatal("There was an error")
			break
		}

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Fatal("Cant unmarshal the payload")
			break
		}

		client.socket.RouteHandler(&request, client)

		log.Println(string(payload))
	}
}

// For writing the messages;
func (client *Client) WriteMessages() {
	defer func() {
		client.socket.RemoveClient(client)
	}()

	for {
		select {
		case message, ok := <-client.messagesChn:
			if !ok {
				if err := client.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Fatal("Failed yto write the close message to client")
				}
				return
			}

			if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Fatal("Failed to write the text message to the user")
				return
			}

		}
	}
}
