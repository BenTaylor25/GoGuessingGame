package main

import "fmt"

func question(question string, answer string) int {
	var guess string
	fmt.Print(question + "\n--> ")
	fmt.Scan(&guess)

	if guess == answer {
		fmt.Print("Correct\n\n")
		return 1
	} else {
		fmt.Printf("Incorrect (%v)\n\n", answer)
		return 0
	}
}

func main() {
	fmt.Print("Quiz\n\n")

	score := 0
	score += question("What's 5 + 5?", "10")
	score += question("What's 5 - 5?", "0")

	fmt.Printf("%v/2", score)
}
