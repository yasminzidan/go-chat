package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type ChatServer struct {
	Messages []string
}

func (s *ChatServer) SendMessage(msg string, reply *[]string) error {

	s.Messages = append(s.Messages, msg)
	*reply = s.Messages
	return nil
}

func main() {
	chatServer := new(ChatServer)
	err := rpc.Register(chatServer)
	if err != nil {
		fmt.Println("Error registering ChatServer:", err)
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()
	fmt.Println("Server listening on port 1234")

	rpc.Accept(listener)
}
