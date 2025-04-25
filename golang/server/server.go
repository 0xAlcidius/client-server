package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Fout bij starten van server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server luistert op poort 12345...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Fout bij accepteren van verbinding:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client verbonden:", conn.RemoteAddr())

	clientReader := bufio.NewReader(conn)
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		message, err := clientReader.ReadString('\n')
		if err != nil {
			fmt.Println("Client heeft verbinding verbroken.")
			break
		}
		fmt.Print("Client zegt:", message)

		fmt.Print("Antwoord: ")
		reply, _ := stdinReader.ReadString('\n')
		_, err = conn.Write([]byte(reply))
		if err != nil {
			fmt.Println("Fout bij verzenden van antwoord:", err)
			break
		}
	}
}
