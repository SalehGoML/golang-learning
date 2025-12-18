package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("command", os.Args[0])

	message := "default message"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	connection, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatalln("can't dial the given address", err)
	}
	fmt.Println("local address", connection.LocalAddr())

	numberOfWrittenByte, wErr := connection.Write([]byte(message))

	if wErr != nil {
		log.Fatalln("can't write data to connection", wErr)
	}

	fmt.Println("numberOfWrittenBytes", numberOfWrittenByte)

	var data = make([]byte, 1024)
	_, rErr := connection.Read(data)

	if rErr != nil {
		log.Fatalln("can't read data from connection", rErr)
	}

	fmt.Println("server response:", string(data))
}
