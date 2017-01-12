//go:generate go-bindata -pkg static -ignore static/static.gen.go -o static/static.gen.go static/...
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/wyattjoh/multi/cmd/multi-client/static"
)

func setupWebServer(addr string, api *API) error {

	wsh := &WebsocketHandler{
		API: api,
	}

	http.Handle("/ws", wsh)
	http.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.Host)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Copy the index file to the response writer.
		buf := bytes.NewReader(static.MustAsset("static/index.html"))

		io.Copy(w, buf)
	})

	return http.ListenAndServe(addr, nil)
}

func main() {
	var addr = flag.String("addr", ":8080", "address for the http/websocket server")
	var apiaddr = flag.String("apiaddr", "127.0.0.1:7070", "address for the api server")

	flag.Parse()

	var wg sync.WaitGroup

	wg.Add(1)

	var api = NewAPI()

	go func() {

		// Setup the web server.
		if err := setupWebServer(*addr, api); err != nil {
			log.Printf("Failed to start web server: %s\n", err.Error())
		}

		// Finish the waitgroup.
		wg.Done()
	}()

	// Start up the api pump.
	if err := api.ListenAndServe(*apiaddr); err != nil {
		log.Printf("Failed to run API pump: %s\n", err.Error())
	}

	wg.Wait()
}
