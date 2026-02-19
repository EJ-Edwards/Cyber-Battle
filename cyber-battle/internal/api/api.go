package api

import (
	"cyber-battle/internal/player"
	"cyber-battle/internal/server"
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
	mux.HandleFunc("/create-room", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	return mux
}

func NewAPI() *API {
	return &API{
		Players: make(map[string]*player.Player),
		Rooms:   make(map[string]*server.Room),
	}
}
