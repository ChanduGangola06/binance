package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	MessageType int
	Data        []byte
}

func main() {
	// Connecting to the remote ws
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	fmt.Printf("Connecting to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer conn.Close()

	// channnels for managing message
	send := make(chan Message)
	done := make(chan struct{})

	// Goroutine for reading mesages
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}
			fmt.Printf("Received: %s\n", message)
		}
	}()

	// Goroutine for sending mesages
	go func() {
		for {
			select {
			case msg := <- send:
				// write that to the websocket connection
				err := conn.WriteMessage(msg.MessageType, msg.Data) {
					if err != nil {

					}
				}
			}()
		}
	}()

	// Read input from the terminal and send it to the websocket server
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type something.../n")
	for scanner.Scan() {
		text := scanner.Text()
		// Send the text to the channnel
		send <- Message{websocket.TextMessage, []byte(text)}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Scanner err: ", err)
	}

}
