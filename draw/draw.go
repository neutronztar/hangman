package draw

import (
	"fmt"
	"strings"
)

// Draw secret phrase with underscores in place of unguessed letters
func DrawSecretPhrase(secretPhrase string, guessedLetters string) {
	for _, v := range secretPhrase {

		if strings.Contains(guessedLetters, string(v)) {
			// Draw letter if user has already guessed it
			fmt.Print(string(v))

		} else {
			// Draw underscore if user hasn't already guessed it
			if string(v) != " " {
				fmt.Print("_")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Print(" ")
	}
	fmt.Println()
}

// Draw the hangman with how many limbs he has left
func DrawMan(limbs int) {
	switch limbs {
	case 0:
		fmt.Println("    +---+")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
	case 1:
		fmt.Println("    +---+")
		fmt.Println("    O   |")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
	case 2:
		fmt.Println("    +---+")
		fmt.Println("    O   |")
		fmt.Println("    |   |")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
	case 3:
		fmt.Println("    +---+")
		fmt.Println("    O   |")
		fmt.Println("    |\\  |")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
	case 4:
		fmt.Println("    +---+")
		fmt.Println("    O   |")
		fmt.Println("   /|\\  |")
		fmt.Println("        |")
		fmt.Println("        |")
		fmt.Println("        |")
	case 5:
		fmt.Println("    +---+")
		fmt.Println("    O   |")
		fmt.Println("   /|\\  |")
		fmt.Println("    |   |")
		fmt.Println("        |")
		fmt.Println("        |")
	case 6:
		fmt.Println("    +---+")
		fmt.Println("    O   |")
		fmt.Println("   /|\\  |")
		fmt.Println("    |   |")
		fmt.Println("     \\  |")
		fmt.Println("        |")
	case 7:
		fmt.Println("    +---+")
		fmt.Println("    O   |")
		fmt.Println("   /|\\  |")
		fmt.Println("    |   |")
		fmt.Println("   / \\  |")
		fmt.Println("        |")
	}
}

func DrawBlankLines(numLines int) {
	for i := 0; i < numLines; i++ {
		fmt.Println()
	}
}
