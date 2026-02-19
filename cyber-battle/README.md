
# Cyber-Battle

A CLI-based multiplayer cybersecurity game written in Go.

## Features
- Multiplayer rooms with custom PINs
- Cybersecurity puzzles and challenges
- Team-based gameplay (Red vs Blue)
- Realistic hacking scenarios
- Learn and compete on the leaderboard

## Requirements
- Go 1.21 or later
- Internet connection (for update checks)

## Installation
1. Clone the repository:
	```sh
	git clone https://github.com/ej-edwards/Cyber-Battle.git
	```
2. Change into the project directory:
	```sh
	cd Cyber-Battle/cyber-battle
	```
3. Run the game:
	```sh
	go run main.go
	```

## Usage
1. When prompted, agree to the rules to start playing.
2. Choose to create a room (set your own PIN) or join an existing room using a PIN.
3. Solve puzzles, use your hacking skills, and compete with others!

## Project Structure
- `main.go`: Entry point
- `internal/server`: Server logic
- `internal/game`: Game logic and state
- `internal/player`: Player logic
- `internal/websites`: Website simulation
- `internal/utils`: Utility functions
- `puzzles/`: Game puzzles
- `tests/`: Tests

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
MIT
