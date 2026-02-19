package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Puzzles struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"`
	Hints       []string `json:"hints"`
	Solution    string   `json:"solution"`
}

// Game represents the main game logic.
type Game struct {
	Puzzles []Puzzles
}

func NewGame() *Game {
	return &Game{}
}

func LoadPuzzlesFromDir(dir string) ([]Puzzles, error) {
	var puzzles []Puzzles
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			var p Puzzles
			if err := json.Unmarshal(data, &p); err != nil {
				return nil, err
			}
			puzzles = append(puzzles, p)
		}
	}
	return puzzles, nil
}

func (g *Game) LoadPuzzles() {
	puzzles, err := LoadPuzzlesFromDir("puzzles")
	if err != nil {
		fmt.Println("Error loading puzzles:", err)
		return
	}
	g.Puzzles = puzzles
	fmt.Printf("Loaded %d puzzles.\n", len(g.Puzzles))
}

func (g *Game) Start() {
	fmt.Println("Game started with", len(g.Puzzles), "puzzles.")
}

func (g *Game) End() {
	fmt.Println("Game ended thanks for playing.")
}
