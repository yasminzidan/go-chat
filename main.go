package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "host.docker.internal:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("connect to chatroom, enter your message: ")
	fmt.Println("Type 'exit' to leave the chatroom.\n")

	for {
		fmt.Print("> ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if msg == "exit" {
			fmt.Println("Exiting chat.")
			break
		}

		msg = fmt.Sprintf("%s: %s", name, msg) // merging the name with the message
		var chatReply []string                 // to hold the chat history
		err = client.Call("ChatServer.SendMessage", msg, &chatReply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
		fmt.Println("\n--- chat history ---")
		for _, m := range chatReply {
			fmt.Println(m)
		}
		fmt.Println("-----------")
	}
}

