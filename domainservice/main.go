//Domain service registers domain names
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/go-nats"
)

//NATSHost - the connection string for NATS
var NATSHost = "https://localhost:4222"

//Domains - Holds domain names
var Domains = make(map[uint64]map[string]string)

var c chan bool

//Domain service
func main() {
	wg := sync.WaitGroup{}

	go func() {
		createsubdomain(c)
	}()
	if <-c == true {
		wg.Done()
	}

}

//failOnError is  single place to handle errors to reduces number of keystrokes per each error handling call.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}

//func createsubdomain(wg *sync.WaitGroup) {
func createsubdomain(c chan bool) {
	//var message []byte
	nc, err := nats.Connect(NATSHost)
	failOnError(err, "Could not connect to NATSHost")
	fmt.Println("Connected to NATSHost. Waiting for subscription...")
	counter := 0

	for {
		nc.Subscribe("createsubdomain", func(m *nats.Msg) {
			//log.Printf("[Order] %s", string(m.Data))
			processSubDomain(m)
		})

		time.Sleep(time.Second)
		counter++
		if counter > 0 {
			c <- true
		}
	}
}

func processSubDomain(m *nats.Msg) {
	//data := m.Data
	//subject := m.Subject
	jdata, err := json.Marshal(string(m.Data))
	failOnError(err, "Failed to marshal to Json")
	
	fmt.Printf("We got some data: %s \n", string(jdata))

}

//Check subdomain
func checkSubDomain(name string) bool {

	return false
}
