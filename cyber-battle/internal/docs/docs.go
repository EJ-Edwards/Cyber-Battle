package docs

import "fmt"

func main() {
	fmt.Println("Welcome to Cyber-Battle!")
}

func Rules() {
	fmt.Println("1. Solve puzzles to progress through the game.")
	fmt.Println("2. Use your hacking skills to overcome challenges.")
	fmt.Println("3. Compete against other players to climb the leaderboard.")
	fmt.Println("4. No cheating or inappropriate behavior.")
	fmt.Println("5. Have fun and learn about cybersecurity!")
	fmt.Println("6. Respect other players and their privacy.")
	fmt.Println("7. Report bugs or vulnerabilities responsibly.")
}

func AgreeToRules() bool {
	var response string
	fmt.Print("Do you agree to the rules? (yes/no): ")
	fmt.Scanln(&response)
	if response == "yes" {
		fmt.Println("Thank you for agreeing! Let's get started!")
		return true
	}
	fmt.Println("You must agree to the rules to play the game. Goodbye!")
	return false
}
