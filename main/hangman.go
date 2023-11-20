package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func wordToGuess(word string, found []string) bool {
	for _, letter := range word {
		if !contains(found, string(letter)) {
			return false
		}
	}
	return true
}

func contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func getRandomWordFromList(words []string) string {
	randIndex := rand.Intn(len(words))
	return words[randIndex]
}

func getUserGuess() string {
	var guess string
	fmt.Print("Entrez une lettre : ")
	fmt.Scanln(&guess)
	return guess
}

func displayGameStatus(word string, found []string, attemptsRemaining int, hangmanStages []string, currentStage int) {
	fmt.Println("\nMot à deviner :")
	displayWord(word, found)
	fmt.Println("\nEssais restants :", attemptsRemaining)
	fmt.Println(hangmanStages[currentStage])
}

func displayWord(word string, found []string) {
	for _, letter := range word {
		if contains(found, string(letter)) {
			fmt.Print(string(letter), " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}
