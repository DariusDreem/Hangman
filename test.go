package main

import (
	"fmt"
	"os"
)

func verifie(liste []int, nbr int) bool {
	for i := 0; i < len(liste); i++ {
		if liste[i] == nbr {
			return false
		}
	}
	return true
}

func pendu(nbr, position int) int {
	jose, _ := os.ReadFile("hangman.txt")
	position += 71 * nbr
	if position >= 710 {
		position = 709
	}
	fmt.Print(string(jose[position-70 : position]))
	return position
}

func verification(word, choice string) []int {
	var listeIndication []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[0] {

		}
	}
	return listeIndication
}
