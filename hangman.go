package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	position := 0
	var choice string
	var letter rune
	LetterFind := ""
	worldbase := motrandom(os.Args[1])
	attemps := 10
	println(worldbase)
	println(len(worldbase))
	println("Good Luck, you have ", attemps, " attempts.")
	showWord := creadumot(worldbase)
	echec := false
	for attemps > 0 {
		for i := 0; i < len(worldbase); i++ {
			print(showWord[i] + " ")
		}
		print("\n" + "\n" + "Choose: ")
		fmt.Scanln(&choice)
		var listeind []int
		for i := 0; i < len(worldbase); i++ {
			if choice[0] == worldbase[i] {
				listeind = append(listeind, i)
			}
		}
		if len(choice) == 1 {
			letter = rune(choice[0])
			if len(verif(LetterFind, choice)) > 0 {
				attemps--
				position = gallows(1, position)
				println("\nalready present in the word,", attemps, "attempts remaining\n")
				echec = true
			}
			LetterFind += choice
			if len(verif(worldbase, choice)) >= 1 {
				index := verif(worldbase, choice)
				for i := 0; i < len(index); i++ {
					showWord[index[i]] = string(letter - 32)
				}
			} else {
				if !echec {
					attemps--
					position = gallows(1, position)
					println("\nNot present in the word,", attemps, "attempts remaining\n")
					echec = false
				}
			}
		} else {
			if choice == worldbase {
				println("\nCongrats !")
				return
			} else {
				attemps -= 2
				position = gallows(2, position)
				println("\nlie! is not the real word,", attemps, "attempts remaining\n")
			}
		}
	}
	println("t'es nul c'était :", worldbase)
}

func motrandom(mot string) string {
	file, _ := os.Open(mot)
	var mots []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		mots = append(mots, fileScanner.Text())
	}
	return mots[rand.Intn(len(mot))]
}

func verif(word, choice string) []int {
	var listeInd []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[i] {
			listeInd = append(listeInd, i)
		}
	}
	return listeInd
}

func creadumot(mot string) []string {
	var mot_cache []string
	nbrlettre := len(mot)/2 - 1
	for i := 0; i < len(mot); i++ {
		mot_cache = append(mot_cache, "_")
	}
	for x := 0; x < nbrlettre; x++ {
		ind := rand.Intn(len(mot))
		mot_cache[ind] = string(mot[ind])
	}
	return mot_cache
}

func gallows(nbr, position int) int {
	jose, _ := os.ReadFile("hangman.txt")
	position += 71 * nbr
	if position >= 710 {
		position = 709
	}
	fmt.Print(string(jose[position-71 : position]))
	return position
}
