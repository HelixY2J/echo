package publisher

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	nc *nats.Conn
}

func InitNATS() (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
		return nil, err
	}
	log.Println("Connected to NATS successfully")
	return nc, nil
}

func NewPublisher(nc *nats.Conn) *Publisher {
	return &Publisher{nc: nc}
}

// func InitNATS() error {
// 	// Connect to NATS server

// 	// db.InitDB()
// 	// defer db.GetDB().Close()

// 	nc, err := nats.Connect(nats.DefaultURL)
// 	if err != nil {
// 		log.Fatalf("Error connecting to NATS server: %v", err)
// 		return err
// 	}
// 	log.Println("Connected to NATS successfully")
// 	defer nc.Close()
// 	return nil
// 	// channels := []string{"task-tracker", "erp"}
// 	// for _, channel := range channels {
// 	// 	go func(ch string) {
// 	// 		err := db.StoreNotification(ch, fmt.Sprintf("New message in %s channel", ch))
// 	// 		if err != nil {
// 	// 			log.Printf("Error storing notification for %s: %v", ch, err)
// 	// 		}
// 	// 		err = nc.Publish(ch, []byte(fmt.Sprintf("New notification for %s", ch)))
// 	// 		if err != nil {
// 	// 			log.Printf("Error publishing to %s: %v", ch, err)
// 	// 		}
// 	// 	}(channel)
// 	// }

// 	//select {}
// }

func (p *Publisher) PublishNotification(data []byte) error {
	if p.nc == nil || !p.nc.IsConnected() {
		log.Printf("NATS connection is not established")
	}

	log.Println("Attempting to publish to natss...")
	if err := p.nc.Publish("notifications", data); err != nil {
		log.Printf("OOps cant publish message: %v\n", err)
		return err
	}
	log.Printf("Published message to NATS: %s\n", string(data))

	return nil
}

// for i := 1; ; i++ {
// 	msg := fmt.Sprintf("Notification %d", i)
// 	err := nc.Publish("notifications", []byte(msg))
// 	if err != nil {
// 		log.Fatalf("Error publishing message: %v", err)
// 	}
// 	fmt.Printf("Published: %s\n", msg)
// 	time.Sleep(3 * time.Second) // Wait for 2 seconds before sending the next message
// }
