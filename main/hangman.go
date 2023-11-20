package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var maxAttempts = 6

func isWordGuessed(word string, found []string) bool {
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
	fmt.Print("Entrer une lettre ")
	fmt.Scanln(&guess)
	return guess
}

func displayGameStatus(word string, found []string, attemptsRemaining int, stage []string) {
	fmt.Println("\nmot a trouver")
	displayWord(word, found)
	fmt.Println("\nessais restant", attemptsRemaining)
	fmt.Println(stage[maxAttempts-attemptsRemaining])
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
