package app

import (
	"fmt"
	"log"

	pb "github.com/Taiters/monstermaker/proto"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type RoomState int

const (
	Lobby RoomState = iota
	Drawing
	Viewing
)

type Room struct {
	Id        string
	Clients   map[string]*Client
	RoomState RoomState
	join      chan *Client
	leave     chan *Client
	update    chan *Client
}

func NewRoom(id string) *Room {
	return &Room{
		Id:        id,
		Clients:   make(map[string]*Client),
		RoomState: Lobby,
		join:      make(chan *Client),
		leave:     make(chan *Client),
		update:    make(chan *Client),
	}
}

func (r *Room) broadcast(message *pb.Event, from *Client) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling proto: %w", err)
	}

	preparedMessage, err := websocket.NewPreparedMessage(websocket.BinaryMessage, data)
	if err != nil {
		return fmt.Errorf("error preparing message: %w", err)
	}

	for _, client := range r.Clients {
		if from != nil && client.Id == from.Id {
			continue
		}

		client.send <- &PreparedClientMessage{preparedMessage}
	}
	return nil
}

func (r *Room) send(message *pb.Event, to *Client) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling proto: %w", err)
	}

	to.send <- &ByteClientMessage{data}
	return nil
}

func (r *Room) onClientJoined(client *Client) {
	r.Clients[client.Id] = client
	if err := r.send(serverWelcome(client, r), client); err != nil {
		log.Printf("Error sending room %s snapshot to %s: %s", r.Id, client.Id, err)
	}
	if err := r.broadcast(clientJoined(client), client); err != nil {
		log.Printf("Error broadcasting client %s joining %s: %s", client.Id, r.Id, err)
	}
	log.Printf("Client %s has joined room %s", client.Id, r.Id)
}

func (r *Room) onClientLeft(client *Client, stop chan *Room) {
	delete(r.Clients, client.Id)
	close(client.send)
	if err := r.broadcast(clientLeft(client), nil); err != nil {
		log.Printf("Error broadcasting client %s leaving %s: %s", client.Id, r.Id, err)
	}
	log.Printf("Client %s has left room %s", client.Id, r.Id)

}

func (r *Room) onClientUpdated(client *Client) {
	if err := r.broadcast(clientUpdated(client), client); err != nil {
		log.Printf("Error broadcasting client update for %s: %s", client.Id, err)
	}
}

func (r *Room) Run(stop chan *Room) {
	log.Printf("Room %s is now running", r.Id)
	defer func() {
		log.Printf("Room %s is closing", r.Id)
		stop <- r
	}()
	for {
		select {
		case client := <-r.join:
			r.onClientJoined(client)
		case client := <-r.leave:
			r.onClientLeft(client, stop)
		case client := <-r.update:
			r.onClientUpdated(client)
		}

		if len(r.Clients) == 0 {
			return
		}
	}
}
