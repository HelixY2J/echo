// subscriber.go
package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	//channel := os.Args[1]

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS server: %v", err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("notifications", func(msg *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to subject: %v", err)
	}

	// Subscribe to the specific channel
	// _, err = nc.Subscribe(channel, func(msg *nats.Msg) {
	// 	fmt.Printf("Received message from %s: %s\n", channel, string(msg.Data))
	// })
	// if err != nil {
	// 	log.Fatalf("Error subscribing to subject: %v", err)
	// }

	// // Keep the connection alive to continue receiving messages
	select {}

}
