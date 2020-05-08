package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println(nats.DefaultURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error:\n %s ", err)
	}

	nc.

	fmt.Println("Published message...")

	for i := 0; i < 100; i++ { // range  time.NewTicker(5 * time.Second).C {
		nc.Publish("customer.updates", []byte(fmt.Sprintf("new message %d", i)))
		time.Sleep(5 * time.Second)
	}

	runtime.Goexit()

}
