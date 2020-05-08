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

	fmt.Println("Published message...")

	for range time.NewTicker(5 * time.Second).C {
		nc.Publish("kpk.pk", []byte("hello world !"))
	}

	runtime.Goexit()

}
