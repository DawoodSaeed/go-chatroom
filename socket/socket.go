package socket

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Socket struct {
	clients ClientMap
	sync.RWMutex
	eventHandlers map[string]EventHandlerFn
}

func NewSocket() *Socket {
	socket := &Socket{
		clients:       make(ClientMap),
		eventHandlers: make(map[string]EventHandlerFn),
	}

	socket.setUpEventHandlers()
	return socket
}

func (s *Socket) setUpEventHandlers() {
	s.eventHandlers[EventMessageType] = sendMessage
}

func sendMessage(e *Event, client *Client) error {
	fmt.Println(e)
	return nil
}

func (s *Socket) RouteHandler(event *Event, client *Client) error {
	if handler, ok := s.eventHandlers[event.Type]; ok {
		if err := handler(event, client); err != nil {
			log.Fatal("Error will executing the event handler")
			return err
		}
		return nil
	} else {
		return errors.New("this Type Of Event Hanlder doesnt exists")
	}
}

func (s *Socket) EstablishSocketConn(w http.ResponseWriter, r *http.Request) {
	log.Println("Go Socket Connection Established;")
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("There is an error with the upgrader")
		return
	}

	newClient := NewClient(conn, s)
	s.AddClient(newClient)
	// You can read as many messages as possible'
	// But u can write message concurrently'
	go newClient.ReadMessages()

	go newClient.WriteMessages()
}

func (s *Socket) AddClient(newClient *Client) {
	s.Lock()
	defer s.Unlock()
	log.Println("Add the client method is called;")
	s.clients[newClient] = true
}

func (s *Socket) RemoveClient(toRemoveClient *Client) {

	s.Lock()
	defer s.Unlock()

	log.Println(" To remove the client method is called;")
	if _, ok := s.clients[toRemoveClient]; ok {
		toRemoveClient.conn.Close()
		delete(s.clients, toRemoveClient)
	}
}
