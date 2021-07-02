package app

import (
	"fmt"
	"log"
	"time"

	pb "github.com/Taiters/monstermaker/proto"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type ClientStatus int

const (
	Joining ClientStatus = iota
	Online
	Offline
)

const (
	writeTimeout     = 10 * time.Second
	heartbeatTimeout = 60 * time.Second
	heartbeatPeriod  = 50 * time.Second
	maxMessageSize   = 256
)

type ClientMessage interface {
	send(c *websocket.Conn) error
}

type PreparedClientMessage struct {
	message *websocket.PreparedMessage
}

func (s *PreparedClientMessage) send(c *websocket.Conn) error {
	return c.WritePreparedMessage(s.message)
}

type ByteClientMessage struct {
	message []byte
}

func (s *ByteClientMessage) send(c *websocket.Conn) error {
	return c.WriteMessage(websocket.BinaryMessage, s.message)
}

type Client struct {
	Id     string
	Name   string
	Status ClientStatus
	conn   *websocket.Conn
	room   *Room
	send   chan ClientMessage
}

func NewClient(id string, conn *websocket.Conn, room *Room) *Client {
	return &Client{
		Id:     id,
		Status: Joining,
		conn:   conn,
		room:   room,
		send:   make(chan ClientMessage),
	}
}

func (c *Client) handleMessage(message *pb.ClientEvent) error {
	switch message.Type {
	case pb.ClientEvent_NAME_CHANGE:
		nameChange := message.GetNameChange()
		if nameChange == nil {
			return fmt.Errorf("no name change data in NAME_CHANGE event")
		}
		c.Name = nameChange.Name
		c.room.update <- c
	}

	return nil
}

func (c *Client) runReader() {
	defer func() {
		c.room.leave <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(heartbeatTimeout))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(heartbeatTimeout))
		return nil
	})
	for {
		messageType, data, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				c.logError("Unexpected close message", err)
			}
			return
		}
		if messageType != websocket.BinaryMessage {
			log.Print("Unexpected text message")
			return
		}
		var message pb.ClientEvent
		if err := proto.Unmarshal(data, &message); err != nil {
			c.logError("Error unmarshalling client event", err)
			return
		}

		if err := c.handleMessage(&message); err != nil {
			c.logError("Error handling message", err)
		}
	}
}

func (c *Client) runWriter() {
	ticker := time.NewTicker(heartbeatPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if ok {
				if err := message.send(c.conn); err != nil {
					c.logError("Error writing message", err)
				}
			} else {
				c.conn.WriteControl(
					websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseNormalClosure, "room closing"),
					time.Now().Add(writeTimeout),
				)
				return
			}
		case <-ticker.C:
			c.conn.WriteControl(
				websocket.PingMessage,
				nil,
				time.Now().Add(writeTimeout),
			)
		}
	}
}

func (c *Client) logError(message string, err error) {
	log.Printf("[Client %s] %s: %s", c.Id, message, err)
}

func (c *Client) Run() {
	go c.runReader()
	go c.runWriter()

	c.room.join <- c
}
