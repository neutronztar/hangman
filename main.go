package main

import (
	"fmt"
	"hangman/draw"
	"hangman/input"
	"strings"
)

var guessedLetters string
var limbs int = 7
var secretPhrase string

func main() {
	draw.DrawBlankLines(200)
	fmt.Println("---HANGMAN---")
	secretPhrase = input.InputSecretPhrase()

	// Main Loop
	for {
		drawGame()

		// User guesses letter
		guess := input.InputLetter(guessedLetters)
		guessedLetters = guessedLetters + guess

		// Lose a limb if your letter isn't in the phrase
		if !strings.Contains(secretPhrase, guess) {
			limbs -= 1
		}

		// Win/lose checks
		if limbs == 0 {
			gameOver()
			return
		}
		if checkForWin() {
			gameWon()
			return
		}
	}
}

func checkForWin() bool {
	for _, v := range secretPhrase {
		if !strings.Contains(guessedLetters+" ", string(v)) {
			// Didn't win yet
			return false
		}
	}
	// Win
	return true
}

func drawGame() {
	draw.DrawBlankLines(100)
	draw.DrawMan(limbs)
	fmt.Println()
	draw.DrawSecretPhrase(secretPhrase, guessedLetters)
	fmt.Println()
	fmt.Println("Guessed Letters:", guessedLetters)
}

func gameOver() {
	draw.DrawBlankLines(100)
	draw.DrawMan(limbs)
	fmt.Println()
	guessedLetters = "abcdefghijklmnopqrstuvwxyz"
	draw.DrawSecretPhrase(secretPhrase, guessedLetters)
	fmt.Println("\nGAME OVER")
}

func gameWon() {
	draw.DrawBlankLines(100)
	draw.DrawMan(limbs)
	fmt.Println()
	draw.DrawSecretPhrase(secretPhrase, guessedLetters)
	fmt.Println("\nYOU WIN")
}
