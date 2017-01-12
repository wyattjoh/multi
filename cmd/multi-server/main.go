package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"

	"github.com/wyattjoh/multi"
)

const (
	pumpTickDuration        = 30 * time.Millisecond
	simulationTickDuration  = 50 * time.Millisecond
	noHeartbeatDeadDuration = 10 * multi.HeartbeatTickDuration
)

var socket *net.UDPConn

type clientState struct {
	id            int64
	addr          *net.UDPAddr
	lastHeartbeat int64
	move          string
}

var server struct {
	state   multi.State
	maxID   int64
	clients map[int64]clientState
	moves   map[int64]string
	sync.RWMutex
}

func pump() error {

	state, err := msgpack.Marshal(server.state)
	if err != nil {
		return err
	}

	for _, client := range server.clients {
		if _, err := socket.WriteToUDP([]byte(state), client.addr); err != nil {
			return err
		}
	}

	return nil
}

func removeClient(id int64) {

	// Decrement the count state
	server.state.Count--

	// Clean up resources.
	delete(server.clients, id)
	delete(server.state.Clients, id)
	delete(server.moves, id)
}

func clearDeadClients() {
	expiry := time.Now().Add(-1 * noHeartbeatDeadDuration).Unix()

	for id, client := range server.clients {
		if client.lastHeartbeat < expiry {
			removeClient(id)

			log.Printf("Client[%d] died: %#v\n", client.id, client.addr.String())
		}
	}
}

func pumper() error {
	simulationTick := time.NewTicker(simulationTickDuration)
	defer simulationTick.Stop()

	pumpTick := time.NewTicker(pumpTickDuration)
	defer pumpTick.Stop()

	heartbeatTick := time.NewTicker(multi.HeartbeatTickDuration)
	defer heartbeatTick.Stop()

	go func() {
		for {

			readbuf := make([]byte, 512)

			// Listen for a new packet, this won't block, but will exit cleanly if
			// there is no data to be read.
			if err := listen(readbuf); err != nil {
				return
			}

		}
	}()

	for {
		select {
		case <-simulationTick.C:

			server.Lock()

			// Simulate the world state.
			server.state.Simulate(server.moves)

			server.Unlock()

		case <-heartbeatTick.C:

			server.Lock()

			// Clear out the dead clients that haven't beaten a heartbeat.
			clearDeadClients()

			server.Unlock()

		case <-pumpTick.C:

			server.RLock()

			// Pump to existing clients.
			if err := pump(); err != nil {
				return err
			}

			server.RUnlock()

		}
	}
}

func handle(msg []byte, newAddr *net.UDPAddr) {
	server.Lock()
	defer server.Unlock()

	// First check to see if we already have this client registered.
	for id, client := range server.clients {
		if client.addr.String() == newAddr.String() {

			// not a new client, increment heartbeat and continue
			client.lastHeartbeat = time.Now().Unix()

			// Parse the message sent down the pipe.
			cmd := string(msg)
			if strings.HasPrefix(cmd, "M;") {
				server.moves[id] = strings.TrimPrefix(cmd, "M;")
			}

			// Update the client.
			server.clients[id] = client

			log.Printf("client[%d] recieved message: %s\n", id, string(msg))

			return
		}
	}

	// A new client!
	log.Printf("new client registered: %#v\n", newAddr.String())

	server.state.Count++
	server.maxID++

	clientID := server.maxID

	server.clients[clientID] = clientState{
		id:            clientID,
		addr:          newAddr,
		lastHeartbeat: time.Now().Unix(),
	}

	server.state.Clients[clientID] = multi.NewClient(clientID)
}

// listen takes a udp connection and listens for incomming requests which
// indicate a new connection, and then a pump routine is dispatched to send to
// that client.
func listen(buf []byte) error {
	n, addr, err := socket.ReadFromUDP(buf)
	if err != nil {
		return err
	}

	if n > 0 {
		handle(buf[0:n], addr)
	}

	return nil
}

func serve(bindAddr string) error {

	// Resolve the UDP address.
	addr, err := net.ResolveUDPAddr("udp4", bindAddr)
	if err != nil {
		return err
	}

	// Create a new socket to listen with.
	socket, err = net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer socket.Close()

	if err := socket.SetWriteBuffer(multi.SendPacketSize); err != nil {
		return err
	}

	// Start the pumper.
	log.Println("pumper started")

	if err := pumper(); err != nil {
		log.Printf("pumper failed: %s\n", err.Error())
		return err
	}

	log.Println("pumper stopped")
	return nil
}

func main() {
	var bindAddr = flag.String("bind", "127.0.0.1:7070", "address to listen on")

	flag.Parse()

	// Setup the initial server.
	server.state = multi.NewState()
	server.clients = make(map[int64]clientState)
	server.moves = make(map[int64]string)

	// Serve the UDP socket.
	if err := serve(*bindAddr); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
