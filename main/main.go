package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Choisissez un fichier :")
	fmt.Println("1. words.txt")
	fmt.Println("2. words2.txt")
	fmt.Println("3. words3.txt")

	var choix int
	fmt.Scanln(&choix)

	var filename string
	switch choix {
	case 1:
		filename = "words.txt"
	case 2:
		filename = "words2.txt"
	case 3:
		filename = "words3.txt"
	default:
		fmt.Println("Choix invalide")
		return
	}

	words, err := readWordsFromFile(filename)
	if err != nil {
		fmt.Println("Fichier non trouvé", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	Word_find := getRandomWordFromList(words)
	letter_Found := make([]string, 0)
	maxAttempts := 6
	Attemps := maxAttempts
	hangmanStages := []string{
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
	Game_Stats := 0

	fmt.Println("HangMan game")
	displayGameStatus(Word_find, letter_Found, Attemps, hangmanStages, Game_Stats)

	for Attemps > 0 {
		guess := getUserGuess()
		letter_Found = append(letter_Found, guess)

		if strings.Contains(Word_find, guess) {
			fmt.Println("Bien jouer")
		} else {
			fmt.Println("Lettre incorrecte, il reste", Attemps, "essais")
			Game_Stats++
			if !wordToGuess(Word_find, letter_Found) {
				Attemps--
			}
		}

		displayGameStatus(Word_find, letter_Found, Attemps, hangmanStages, Game_Stats)

		if wordToGuess(Word_find, letter_Found) {
			fmt.Println("Bravo le mot était :", Word_find)
			break
		}
	}

	if Attemps == 0 {
		fmt.Println("Le mot était :", Word_find)
	}

	if Attemps == 0 && !wordToGuess(Word_find, letter_Found) {
		fmt.Println("Perdu vous n'avez plus d'éssais")
	}
}
