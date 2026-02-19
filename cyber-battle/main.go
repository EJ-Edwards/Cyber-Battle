package main

import (
	"bytes"
	docs "cyber-battle/internal/docs"
	"cyber-battle/internal/updater"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		// Create a room via API
		fmt.Print("Enter room name: ")
		var name string
		fmt.Scanln(&name)
		fmt.Print("Enter max players (2-4): ")
		var maxPlayers int
		fmt.Scanln(&maxPlayers)
		fmt.Print("Set a PIN for your room: ")
		var pin string
		fmt.Scanln(&pin)
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":       name,
			"maxPlayers": maxPlayers,
			"pin":        pin,
		})
		resp, err := http.Post("http://localhost:8080/create-room", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			fmt.Println("Error connecting to API:", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	case 2:
		// Join a room via API
		fmt.Print("Enter room PIN: ")
		var pin string
		fmt.Scanln(&pin)
		fmt.Print("Enter your player name: ")
		var name string
		fmt.Scanln(&name)
		reqBody, _ := json.Marshal(map[string]interface{}{
			"pin":  pin,
			"name": name,
		})
		resp, err := http.Post("http://localhost:8080/join-room", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			fmt.Println("Error connecting to API:", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	case 3:
		fmt.Println("Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice.")
	}
}
