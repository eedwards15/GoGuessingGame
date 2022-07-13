package main

import (
	"GoGuessingGame/Helpers"
	"GoGuessingGame/Systems"
	"GoGuessingGame/models"
	"fmt"
)

type GameEngine struct {
	inputSystem  *Systems.InputSystem
	sceneManager *Systems.SceneManager
	State        models.GameState
}

//Returns a new instance of Game
func NewGame() *GameEngine {
	var game = &GameEngine{}
	game.sceneManager = Systems.InitSceneManager()
	game.inputSystem = Systems.InitInputSystem()
	game.State = models.GameState{}
	game.State.Guesses = 0
	game.State.SecretNumber = Helpers.GenerateRandom(1, 100)
	return game
}

//Game Loop
func (game *GameEngine) Update() bool {
	game.State.Guesses = game.State.Guesses + 1
	fmt.Println("Please input your guess")

	guess, err := game.inputSystem.GuessValue()
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return false
	}

	if game.sceneManager.GuessMessages(guess, game.State.SecretNumber) {
		return true
	}

	return false
}
