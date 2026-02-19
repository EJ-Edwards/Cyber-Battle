package api

import (
	"cyber-battle/internal/player"
	"cyber-battle/internal/server"
	"encoding/json"
	"net/http"
)

type API struct {
	Players map[string]*player.Player
	Rooms   map[string]*server.Room
}

type CreateRoomRequest struct {
	Name       string
	MaxPlayers int
	Pin        string
}

type newPlayerrequest struct {
	Name string
	id   string
}

func NetHTTPHandler(api *API) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Cyber-Battle API!"))
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/create-room", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req CreateRoomRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.MaxPlayers < 2 || req.MaxPlayers > 4 {
			req.MaxPlayers = 4
		}
		if req.Name == "" || req.Pin == "" {
			http.Error(w, "Room name and PIN required", http.StatusBadRequest)
			return
		}
		if _, exists := api.Rooms[req.Name]; exists {
			http.Error(w, "Room name already exists", http.StatusConflict)
			return
		}
		room := &server.Room{
			Name:       req.Name,
			Pin:        req.Pin,
			Clients:    []server.Client{},
			MaxPlayers: req.MaxPlayers,
		}
		api.Rooms[req.Name] = room
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"Room created","pin":"` + req.Pin + `"}`))
	})

	mux.HandleFunc("/join-room", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Pin  string `json:"pin"`
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		var foundRoom *server.Room
		for _, room := range api.Rooms {
			if room.Pin == req.Pin {
				foundRoom = room
				break
			}
		}
		if foundRoom == nil {
			http.Error(w, "No room found with that PIN", http.StatusNotFound)
			return
		}
		if len(foundRoom.Clients) >= foundRoom.MaxPlayers {
			http.Error(w, "Room is full", http.StatusForbidden)
			return
		}
		foundRoom.Clients = append(foundRoom.Clients, server.Client{Name: req.Name})
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Joined room successfully"}`))
	})
	return mux
}

func NewAPI() *API {
	return &API{
		Players: make(map[string]*player.Player),
		Rooms:   make(map[string]*server.Room),
	}
}
