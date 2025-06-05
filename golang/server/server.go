package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on port 1234...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	clientReader := bufio.NewReader(conn)
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		message, err := clientReader.ReadString('\n')
		if err != nil {
			fmt.Println("Client has disconnected.")
			break
		}
		fmt.Print("Client says:", message)

		fmt.Print("Reply: ")
		reply, _ := stdinReader.ReadString('\n')
		_, err = conn.Write([]byte(reply))
		if err != nil {
			fmt.Println("Error sending reply:", err)
			break
		}
	}
}
