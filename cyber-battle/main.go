package main

import (
	docs "cyber-battle/internal/docs"
	"cyber-battle/internal/server"
	"cyber-battle/internal/updater"
	"fmt"
	"os"
)

func newmain() {
	updater.NotifyIfUpdate()
	docs.Rules()
	if !docs.AgreeToRules() {
		fmt.Println("You must agree to the rules to play the game. Goodbye!")
		os.Exit(0)
	}
	menu()
}

func menu() {
	fmt.Println("\nChoose an option:")
	fmt.Println("1. Create a room")
	fmt.Println("2. Join a public room")
	fmt.Println("3. Exit")
	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		server.CreateRoom()
	case 2:
		fmt.Print("Enter room PIN: ")
		var pin string
		fmt.Scanln(&pin)
		fmt.Print("Enter your player name: ")
		var name string
		fmt.Scanln(&name)
		joined := server.JoinRoom(pin, server.Client{Name: name})
		if !joined {
			fmt.Println("Failed to join room. Check the PIN or if the room is full.")
		} else {
			fmt.Println("Successfully joined the room!")
		}
	case 3:
		fmt.Println("Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice.")
	}
}
