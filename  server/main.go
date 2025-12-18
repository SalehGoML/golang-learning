package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		network = "tcp"
		address = "127.0.0.1:8080"
	)

	// create new listener
	listener, err := net.Listen(network, address)

	if err != nil {
		log.Fatalln("cant't listen or given address:", address, err)
	}

	fmt.Println("can't listen on", listener.Addr())

	//time.Sleep(5 * time.Second)

	for {
		// listen for new connections
		connection, aErr := listener.Accept()
		if aErr != nil {
			log.Println("can't listen to new connection", aErr)

			continue
		}

		//fmt.Println("client address", connection.RemoteAddr())

		// process request
		var data = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(data)
		if rErr != nil {
			log.Println("can't read data from connection", rErr)

			continue
		}

		//fmt.Println("numberOfReadBytes", numberOfReadBytes)
		//fmt.Println("- received data", string(data))
		fmt.Printf("client address: %s, numberOfReadBytes: %d, data: %s\n", connection.RemoteAddr(), numberOfReadBytes, string(data))

		_, wErr := connection.Write([]byte(`your message received`))
		if wErr != nil {
			log.Println("can't write data to connection", wErr)

			continue
		}
	}
}
