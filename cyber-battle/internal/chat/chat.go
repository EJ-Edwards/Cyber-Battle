package chat

import (
	"fmt"
)

type Message struct {
	Sender  string
	Team    string // "red", "blue", or "all"
	Content string
}

// SendMessage sends a message to the appropriate team or all players.
func SendMessage(msg Message, currentTeam string) {
	if msg.Team == "all" || msg.Team == currentTeam {
		switch msg.Team {
		case "red":
			fmt.Println("[Red Team]", msg.Sender+":", msg.Content)
		case "blue":
			fmt.Println("[Blue Team]", msg.Sender+":", msg.Content)
		default:
			fmt.Println("[All]", msg.Sender+":", msg.Content)
		}
	}
}

// TeamChat allows a player to send a message only to their team.
func TeamChat(sender, team, content string) {
	SendMessage(Message{Sender: sender, Team: team, Content: content}, team)
}

// GlobalChat allows a player to send a message to all players.
func GlobalChat(sender, content string) {
	SendMessage(Message{Sender: sender, Team: "all", Content: content}, "all")
}

func ReceiveMessage() Message {
	// Placeholder for receiving a message
	return Message{Sender: "System", Team: "all", Content: "Received a message"}
}
