package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println(nats.DefaultURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error:\n %s ", err)
	}

	fmt.Println("Receiving message...")

	nc.Subscribe("kpk.pk", func(m *nats.Msg) {
		log.Printf("[Received kpk.pk] %s", string(m.Data))
	})

	nc.Subscribe("kpk.*", func(m *nats.Msg) {
		log.Printf("[Received kpk.*] %s", string(m.Data))
	})

	runtime.Goexit()
}
