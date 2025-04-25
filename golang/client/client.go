package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./client <server_ip:port>")
		return
	}

	address := os.Args[1]
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server", address)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Message: ")
		msg, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error sending:", err)
			break
		}

		reply := make([]byte, 1024)
		n, err := conn.Read(reply)
		if err != nil {
			fmt.Println("Connection closed.")
			break
		}
		fmt.Println("Server answer:", string(reply[:n]))
	}
}
