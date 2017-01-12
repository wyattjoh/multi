package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebsocketHandler wraps the API and the websocket connections.
type WebsocketHandler struct {
	API *API
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (wsh *WebsocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket: connection upgrading failed: %s\n", err.Error())
		return
	}
	defer con.Close()

	log.Println("websocket: connection established")

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			// Remarshal the state as the payload to send back to the client
			message, err := json.Marshal(wsh.API.State())
			if err != nil {
				log.Printf("websocket: marshaling failed: %s\n", err.Error())
				return
			}

			err = con.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("websocket: write failed: %s\n", err.Error())
				return
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()

		for {
			mt, message, err := con.ReadMessage()
			if err != nil {
				log.Printf("websocket: read failed: %s\n", err.Error())
				return
			}

			if mt == websocket.TextMessage {

				// Send the message back to the server.
				wsh.API.Write(message)
			}
		}
	}()

	wg.Wait()

	log.Println("websocket: connection closed")
}
