package main

import "fmt"

func main() {
	var game = NewGame()
	fmt.Println("Guess a number between 1 and 100")
	for {
		if game.Update() {
			break
		}
	}
}
