package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the Hangman Game!")
	fmt.Println("Choose a file:")
	fmt.Println("1. Words 1")
	fmt.Println("2. Words 2")
	fmt.Println("3. Words 3")

	var choice int
	fmt.Scanln(&choice)

	var filename string
	switch choice {
	case 1:
		filename = "words.txt"
	case 2:
		filename = "words2.txt"
	case 3:
		filename = "words3.txt"
	default:
		fmt.Println("fichier introuvable")
		return
	}

	words, err := readWordsFromFile(filename)
	if err != nil {
		fmt.Println("fichier introuvable", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	word := getRandomWordFromList(words)
	letter := make([]string, 0)
	maxAttempts := 6
	attempts := maxAttempts
	stage := []string{
		`
  +---+
  |   |
      |
      |
      |
      |
=======
`,
		`
  +---+
  |   |
  O   |
      |
      |
      |
=======
`,
		`
  +---+
  |   |
  O   |
  |   |
      |
      |
=======
`,
		`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=======
`,
		`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=======
`,
		`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=======
`,
		`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=======
`,
	}

	fmt.Println("Hangman")
	displayGameStatus(word, letter, attempts, stage)

	for attempts > 0 {
		guess := getUserGuess()
		letter = append(letter, guess)

		guessedCorrectly := isWordGuessed(word, letter)

		if strings.Contains(word, guess) {
			fmt.Println("bravo")
		} else {
			fmt.Println("lettre incorrecte il reste", attempts, "essais")
			if !guessedCorrectly {
				attempts--
			}
		}

		displayGameStatus(word, letter, attempts, stage)

		if guessedCorrectly {
			fmt.Println("Bravo vous avez trouve le mot", word)
			break
		}
	}

	if attempts == 0 {
		fmt.Println("Dommage le mot etait", word)
	}
}
