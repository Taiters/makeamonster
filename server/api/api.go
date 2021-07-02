package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/Taiters/monstermaker/app"
	"github.com/Taiters/monstermaker/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	Rooms    map[string]*app.Room
	upgrader *websocket.Upgrader
	roomLock sync.RWMutex
	remove   chan *app.Room
}

func newServer(upgrader *websocket.Upgrader) *Server {
	return &Server{
		Rooms:    make(map[string]*app.Room),
		upgrader: upgrader,
		remove:   make(chan *app.Room),
	}
}

func (server *Server) run() {
	for {
		select {
		case room := <-server.remove:
			server.removeRoom(room)
			log.Printf("Room %s removed from server", room.Id)
		}
	}
}

func (server *Server) createRoom() *app.Room {
	server.roomLock.Lock()
	defer server.roomLock.Unlock()

	var availableId string
	for {
		availableId = utils.GenerateString(5)
		if _, exists := server.Rooms[availableId]; !exists {
			break
		}
	}

	room := app.NewRoom(availableId)
	server.Rooms[availableId] = room
	return room
}

func (server *Server) removeRoom(room *app.Room) {
	server.roomLock.Lock()
	defer server.roomLock.Unlock()

	delete(server.Rooms, room.Id)
}

func (server *Server) getRoom(id string) (*app.Room, bool) {
	server.roomLock.RLock()
	defer server.roomLock.RUnlock()

	room, exists := server.Rooms[id]
	return room, exists
}

func (server *Server) create(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection during create: %s", err)
		return
	}

	room := server.createRoom()
	client := app.NewClient(uuid.NewString(), conn, room)

	go room.Run(server.remove)
	go client.Run()
}

func (server *Server) status(w http.ResponseWriter, r *http.Request) {
	server.roomLock.RLock()
	defer server.roomLock.RUnlock()

	json.NewEncoder(w).Encode(server)
}

func (server *Server) join(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if room, exists := server.getRoom(vars["roomId"]); exists {
		conn, err := server.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print(err)
			http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
			return
		}

		client := app.NewClient(uuid.NewString(), conn, room)
		go client.Run()
	} else {
		http.Error(w, "Room not found", http.StatusNotFound)
	}
}

func BootstrapServer(router *mux.Router) {
	upgrader := websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	server := newServer(&upgrader)

	router.HandleFunc("/create", server.create)
	router.HandleFunc("/join/{roomId}", server.join)
	router.HandleFunc("/status", server.status)

	go server.run()
}
