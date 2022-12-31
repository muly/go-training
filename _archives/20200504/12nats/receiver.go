package main

import (
	"fmt"
	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

func main() {
	fmt.Println(nats.DefaultURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error:\n %s ", err)
	}

	fmt.Println("Receiving message...")

	s, err := nc.Subscribe("customer.updates", func(m *nats.Msg) {
		log.Printf("[Received kpk.pk] %s", string(m.Data))

	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// time.Sleep(5 * time.Second)
	// s.Unsubscribe()

	s.AutoUnsubscribe(10)

	// nc.Subscribe("kpk.*", func(m *nats.Msg) {
	// 	log.Printf("[Received kpk.*] %s", string(m.Data))
	// })

	runtime.Goexit()
}
