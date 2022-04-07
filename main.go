package main

import (
	"fmt"
	"hangman/draw"
	"hangman/input"
	"strings"
)

type gameState struct {
	guessedLetters string
	limbs          int
	secretPhrase   string
}

func main() {
	draw.DrawBlankLines(200)
	fmt.Println("---HANGMAN---")

	state := gameState{
		guessedLetters: "",
		limbs:          7,
		secretPhrase:   input.InputSecretPhrase(),
	}

	// Main Loop
	for {
		drawGame(state)

		// User guesses letter
		guess := input.InputLetter(state.guessedLetters)
		state.guessedLetters = state.guessedLetters + guess

		// Lose a limb if your letter isn't in the phrase
		if !strings.Contains(state.secretPhrase, guess) {
			state.limbs -= 1
		}

		// Win/lose checks
		if state.limbs == 0 {
			gameOver(state)
			return
		}
		if checkForWin(state) {
			gameWon(state)
			return
		}
	}
}

func checkForWin(state gameState) bool {
	for _, v := range state.secretPhrase {
		if !strings.Contains(state.guessedLetters+" ", string(v)) {
			// Didn't win yet
			return false
		}
	}
	// Win
	return true
}

func drawGame(state gameState) {
	draw.DrawBlankLines(100)
	draw.DrawMan(state.limbs)
	fmt.Println()
	draw.DrawSecretPhrase(state.secretPhrase, state.guessedLetters)
	fmt.Println()
	if state.guessedLetters != "" {
		fmt.Println("Guessed Letters:", state.guessedLetters)
	}
}

func gameOver(state gameState) {
	draw.DrawBlankLines(100)
	draw.DrawMan(state.limbs)
	fmt.Println()
	state.guessedLetters = "abcdefghijklmnopqrstuvwxyz"
	draw.DrawSecretPhrase(state.secretPhrase, state.guessedLetters)
	fmt.Println("\nGAME OVER")
}

func gameWon(state gameState) {
	draw.DrawBlankLines(100)
	draw.DrawMan(state.limbs)
	fmt.Println()
	draw.DrawSecretPhrase(state.secretPhrase, state.guessedLetters)
	fmt.Println("\nYOU WIN")
}
