package Systems

import (
	"fmt"
)

//Handles displaying outputs to the player.
type SceneManager struct{}

func InitSceneManager() *SceneManager {
	return &SceneManager{}
}

func (sceneManager *SceneManager) GuessMessages(guess int, secretNumber int) bool {
	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number. Try again")
		return false
	} else if guess < secretNumber {
		fmt.Println("Your guess is smaller than the secret number. Try again")
		return false
	} else {
		fmt.Println("Correct, you are a Legend!")
		return true
	}
}
