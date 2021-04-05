package gohack

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//EchoServer ...
/*

 */
type EchoServer struct {
	Address string
	Port    int
}

//MakeServer ...
/*

 */
func MakeServer(address string, port int) EchoServer {
	server := EchoServer{
		Address: address,
		Port:    port,
	}
	return server
}

func echo(connObj net.Conn) {
	defer connObj.Close()
	reader := bufio.NewReader(connObj)
	// bytes := make([]byte, 512)
	s, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("E: ", err)
	}
	log.Printf("Writing data...")

	log.Println("Writing data")
	writer := bufio.NewWriter(connObj)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("E: ", err)
	}
	writer.Flush()
}

//Listen ...
func (server *EchoServer) Listen() {

	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", fmt.Sprintf(server.Address+":%d", server.Port))

	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening... ")
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()

		log.Println("Received connection")

		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}
