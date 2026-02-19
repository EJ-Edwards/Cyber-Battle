package server

import (
	"fmt"
	"time"
)

type Client struct {
	Name string
	Team string // "red" or "blue"
}

type Room struct {
	Name       string
	Pin        string
	Clients    []Client
	MaxPlayers int
}

var rooms = make(map[string]*Room)

func CreateRoom() {
	fmt.Print("Enter room name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Print("Enter max players (2-4): ")
	var maxPlayers int
	fmt.Scanln(&maxPlayers)
	if maxPlayers < 2 || maxPlayers > 4 {
		fmt.Println("Invalid number of players. Defaulting to 4.")
		maxPlayers = 4
	}
	fmt.Print("Set a PIN for your room: ")
	var pin string
	fmt.Scanln(&pin)
	room := &Room{Name: name, Pin: pin, Clients: []Client{}, MaxPlayers: maxPlayers}
	rooms[name] = room
	fmt.Printf("Room '%s' created with PIN: %s (max %d players)\n", name, pin, maxPlayers)
	fmt.Println("Waiting for other players to join...")
	waitForPlayers(room)
}

func waitForPlayers(room *Room) {
	for len(room.Clients) < room.MaxPlayers {
		fmt.Printf("Current players: %d/%d\n", len(room.Clients), room.MaxPlayers)
		fmt.Println("Waiting for more players...")
		time.Sleep(3 * time.Second)
	}
	assignTeams(room)
	fmt.Println("All players joined! Teams assigned:")
	for _, c := range room.Clients {
		fmt.Printf("%s: %s\n", c.Name, c.Team)
	}
}

func assignTeams(room *Room) {
	redCount := room.MaxPlayers / 2
	for i := range room.Clients {
		if i < redCount {
			room.Clients[i].Team = "red"
		} else {
			room.Clients[i].Team = "blue"
		}
	}
}

func Start() {
	fmt.Println("Cyber Battle server started.")
}

func JoinRoom(pin string, client Client) bool {
	for _, room := range rooms {
		if room.Pin == pin {
			if len(room.Clients) >= room.MaxPlayers {
				fmt.Println("Room is full.")
				return false
			}
			room.Clients = append(room.Clients, client)
			fmt.Printf("Client '%s' joined room '%s'\n", client.Name, room.Name)
			return true
		}
	}
	fmt.Printf("No room found with PIN: %s\n", pin)
	return false
}

func Stop() {
	fmt.Println("Cyber Battle server stopped.")
}
