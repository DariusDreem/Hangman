package main

import "os"

func verifie(liste []int, nbr int) bool {
	for i := 0; i < len(liste); i++ {
		if liste[i] == nbr {
			return false
		}
	}
	return true
}

func pendu(nombre, position int) int {
	jose, _ := os.ReadFile("hangman.txt")
}
