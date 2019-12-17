package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Guess a number between 1 and 100")
	secretNumber := generateRandomInteger(1, 100)

	var attempts int

	for {
		attempts = attempts + 1

		fmt.Println("Please input your guess")

		input := GuessValue()
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		if GuessMessages(guess, secretNumber) {
			break
		}

	}
}

func GuessValue() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}

func GuessMessages(guess int, secretNumber int) bool {
	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number. Try again")
		return false
	} else if guess < secretNumber {
		fmt.Println("Your guess is smaller than the secret number. Try again")
		return false
	} else {
		fmt.Println("Correct, you Legend!")
		return true
	}
}

func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
