package main

import (
	"log"
	"net/http"

	"github.com/HelixY2J/echo/common/db"
	handlers "github.com/HelixY2J/echo/common/handler"
	"github.com/HelixY2J/echo/common/publisher"
)

func main() {
	// db
	obj, err := db.InitDB("postgres://pos:pom@localhost:5436/echoDB?sslmode=disable")
	if err != nil {
		log.Fatal("failed to init database")
	}
	defer obj.Close()

	// nats
	nc, err := publisher.InitNATS()
	if err != nil {
		log.Fatal("failed to init nats")
	}
	defer nc.Close()

	// inject
	pub := publisher.NewPublisher(nc)
	store := db.NewStore(obj)

	// http
	http.HandleFunc("/publish", handlers.PublishHandler(pub, store))

	if err := http.ListenAndServe(":8086", nil); err != nil {
		log.Fatalf("srver stopped: %v", err)
	}
	log.Println("server started ")
}
