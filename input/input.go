package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const validPhraseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz"
const validGuessChars = "abcdefghijklmnopqrstuvwxyz"

// Get input phrase from user
func InputSecretPhrase() string {
	for {
		fmt.Print("Type secret phrase: ")

		// User input
		in := bufio.NewReader(os.Stdin)
		userInput, _ := in.ReadString('\n')

		// Get rid of carriage return at the end
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		// Make sure input is valid
		err := ValidatePhrase(userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return userInput
	}
}

// Validate phrase
func ValidatePhrase(userInput string) error {
	// Make sure it only has letters and spaces
	for _, char := range userInput {
		if !strings.Contains(validPhraseChars, string(char)) {
			return errors.New("secret phrase can only contain letters and spaces")
		}
	}

	// Make sure it isn't blank
	if userInput == "" {
		return errors.New("secret phrase cannot be blank")
	}

	// Make sure it isn't too long
	if len(userInput) > 200 {
		return errors.New("secret phrase max length is 200 chars")
	}

	return nil
}

// Get letter guess from user
func InputLetter(guessedLetters string) string {
	for {
		fmt.Print("Guess a letter: ")

		// User input
		in := bufio.NewReader(os.Stdin)
		userInput, _ := in.ReadString('\n')

		// Get rid of carriage return at the end and make it lowercase
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.ToLower(userInput)

		// Make sure input is valid
		err := ValidateLetter(userInput, guessedLetters)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return userInput
	}
}

// Validate letter guess
func ValidateLetter(userInput string, guessedLetters string) error {
	// Make sure length is 1
	if len(userInput) != 1 {
		return errors.New("length of letter guess must be 1")
	}

	// Make sure it is a letter
	if !strings.Contains(validGuessChars, userInput) {
		return errors.New("guess must be a letter")
	}

	// Make sure it hasn't been guessed yet
	if strings.Contains(guessedLetters, userInput) {
		return errors.New("already guessed that letter")
	}

	return nil
}
