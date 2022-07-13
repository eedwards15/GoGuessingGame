package Systems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Handles input from the user.
type InputSystem struct{}

func InitInputSystem() *InputSystem {
	return &InputSystem{}
}

func (inputSystem *InputSystem) GuessValue() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	intVar, err := strconv.Atoi(input)

	return intVar, err
}
