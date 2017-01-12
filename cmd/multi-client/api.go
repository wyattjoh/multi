package main

import (
	"bufio"
	"log"
	"net"
	"sync"
	"time"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"

	"github.com/wyattjoh/multi"
)

// APIState wraps a multi.State with a RWMutex.
type APIState struct {
	multi.State
	sync.RWMutex
}

// NewAPI creates a new instance of API.
func NewAPI() *API {
	return &API{
		writeChan: make(chan []byte, 100),
	}
}

// API pumps commands to the API and recieves state from the API.
type API struct {
	state     APIState
	socket    *net.UDPConn
	writeChan chan []byte
}

// ListenAndServe runs the API pumping service.
func (api *API) ListenAndServe(serverAddr string) error {
	addr, err := net.ResolveUDPAddr("udp4", serverAddr)
	if err != nil {
		return err
	}

	api.socket, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer api.socket.Close()

	var wg sync.WaitGroup

	// Add the reader and writer.
	wg.Add(2)

	go func() {

		defer wg.Done()

		log.Println("api.write: started")

		if err := api.StartWriter(); err != nil {
			log.Printf("api.write: failed: %s\n", err.Error())
			return
		}

		log.Println("api.write: stopped")

	}()

	go func() {

		defer wg.Done()

		log.Println("api.read: started")

		if err := api.StartReader(); err != nil {
			log.Printf("api.read: failed: %s\n", err.Error())
			return
		}

		log.Println("api.read: stopped")

	}()

	// Write the connect message out.
	api.Write([]byte(multi.ConnectMessage))

	// Wait till the reader/writer shutdown.
	wg.Wait()

	return nil
}

func (api *API) Write(data []byte) {
	api.writeChan <- data
}

// State retrieves the state from the API.
func (api *API) State() multi.State {

	api.state.RLock()
	state := api.state.State
	api.state.RUnlock()

	return state
}

// StartWriter starts up the writer to write to the socket.
func (api *API) StartWriter() error {
	heartbeatTimer := time.NewTimer(multi.HeartbeatTickDuration)

	for {
		select {
		case <-heartbeatTimer.C:

			log.Println("api.write: writing heartbeat")

			// Push the heartbeat into the socket.
			if _, err := api.socket.Write([]byte(multi.HeartbeatMessage)); err != nil {
				return err
			}

			// Reset the heartbeat to fire again in the tick duration.
			heartbeatTimer.Reset(multi.HeartbeatTickDuration)

		case data := <-api.writeChan:

			if !heartbeatTimer.Stop() {
				<-heartbeatTimer.C
			}

			log.Println("api.write: writing data")

			// Write the data out to the socket.
			if _, err := api.socket.Write(data); err != nil {
				return err
			}

			heartbeatTimer.Reset(multi.HeartbeatTickDuration)

		}
	}
}

// StartReader reads from the socket and updates the state.
func (api *API) StartReader() error {

	reader := bufio.NewReader(api.socket)

	for {
		var buf = make([]byte, multi.SendPacketSize)
		n, err := reader.Read(buf)
		if err != nil {
			return err
		}

		var newState multi.State
		if err := msgpack.Unmarshal(buf[0:n], &newState); err != nil {
			return err
		}

		api.state.RLock()
		if api.state.Count != newState.Count {
			log.Printf("Server has %d clients connected.\n", newState.Count)
		}
		api.state.RUnlock()

		api.state.Lock()
		api.state.State = newState
		api.state.Unlock()
	}
}
