package multi

import "time"

// SendPacketSize is the size of a packet sent by the server during state
// pushes.
const SendPacketSize = 8192

// HeartbeatTickDuration is the time that we expect clients to send their
// heartbeats on.
const HeartbeatTickDuration = 500 * time.Millisecond

// NewState creates a new State object.
func NewState() State {
	return State{
		Clients: make(map[int64]Client),
		Now:     time.Now(),
	}
}

// State is the representation of the game map.
type State struct {
	Count   int64            `json:"count"`
	Clients map[int64]Client `json:"clients"`
	Now     time.Time        `json:"now"`
}

// Simulate simulates the world state.
func (s *State) Simulate(moves map[int64]string) {
	for clientID, move := range moves {

		client := s.Clients[clientID]

		switch move {
		case "L":
			client.Speed--
		case "R":
			client.Speed++
		case "":
			if client.Speed > 0 {
				client.Speed--
			} else if client.Speed < 0 {
				client.Speed++
			}
		}

		if client.Speed > 20 {
			client.Speed = 20
		} else if client.Speed < -20 {
			client.Speed = -20
		}

		client.Position += client.Speed

		if client.Position < 0 {
			client.Position += 360
		} else if client.Position >= 360 {
			client.Position -= 360
		}

		s.Clients[clientID] = client
	}
}

// NewClient creates a client for which is connected.
func NewClient(id int64) Client {
	return Client{
		ID: id,
	}
}

// Client is a player.
type Client struct {
	ID       int64 `json:"i"`
	Position int64 `json:"p"`
	Speed    int64 `json:"s"`
}

// HeartbeatMessage is sent when the client Heartbeats.
const HeartbeatMessage = "HEARTBEAT"

// ConnectMessage is sent when the client Connects.
const ConnectMessage = "CONNECT"
