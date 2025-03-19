package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	for i := 1; i <= 10; i++ {
		task := fmt.Sprintf("Task #%d", i)
		err := nc.Publish("tasks", []byte(task))
		if err != nil {
			log.Println("Failed to publish:", err)
		} else {
			log.Println("Published:", task)
		}
		time.Sleep(500 * time.Microsecond) 
	}

	log.Println("All tasks published.")
}
