package app

import pb "github.com/Taiters/monstermaker/proto"

func serverWelcome(client *Client, room *Room) *pb.Event {
	return &pb.Event{
		Type: pb.Event_SERVER_WELCOME,
		Data: &pb.Event_Welcome{
			Welcome: &pb.Event_ServerWelcome{
				ClientId: client.Id,
				Room:     mapRoom(room),
			},
		},
	}
}

func clientJoined(client *Client) *pb.Event {
	return &pb.Event{
		Type: pb.Event_CLIENT_JOINED,
		Data: &pb.Event_Client_{
			Client: mapClient(client),
		},
	}
}

func clientLeft(client *Client) *pb.Event {
	return &pb.Event{
		Type: pb.Event_CLIENT_LEFT,
		Data: &pb.Event_Client_{
			Client: mapClient(client),
		},
	}
}

func clientUpdated(client *Client) *pb.Event {
	return &pb.Event{
		Type: pb.Event_CLIENT_UPDATED,
		Data: &pb.Event_Client_{
			Client: mapClient(client),
		},
	}
}

func mapRoom(room *Room) *pb.Event_Room {
	return &pb.Event_Room{
		Id:      room.Id,
		Clients: mapClients(room.Clients),
	}
}

func mapClients(clients map[string]*Client) map[string]*pb.Event_Client {
	clientsMap := make(map[string]*pb.Event_Client)
	for id, client := range clients {
		clientsMap[id] = mapClient(client)
	}
	return clientsMap
}

func mapClient(client *Client) *pb.Event_Client {
	return &pb.Event_Client{
		Id:   client.Id,
		Name: client.Name,
	}
}
