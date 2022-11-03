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
	worldbase := random_word(os.Args[1])
	attemps := 10
	println(worldbase)
	println(len(worldbase))
	println("Good Luck, you have ", attemps, " attempts.")
	showWord := word_choice(worldbase)
	fail := false
	for attemps > 0 {
		for i := 0; i < len(worldbase); i++ {
			print(showWord[i] + " ")
		}
		print("\n" + "\n" + "Choose: ")
		fmt.Scanln(&choice)
		var listInd []int
		for i := 0; i < len(worldbase); i++ {
			if choice[0] == worldbase[i] {
				listInd = append(listInd, i)
			}
		}
		if len(choice) == 1 {
			letter = rune(choice[0])
			if len(verif(LetterFind, choice)) > 0 {
				attemps--
				position = gallows(1, position)
				println("\nalready present in the word,", attemps, "attempts remaining\n")
				fail = true
			}
			LetterFind += choice
			if len(verif(worldbase, choice)) >= 1 {
				index := verif(worldbase, choice)
				for i := 0; i < len(index); i++ {
					showWord[index[i]] = string(letter - 32)
				}
			} else {
				if !fail {
					attemps--
					position = gallows(1, position)
					println("\nNot present in the word,", attemps, "attempts remaining\n")
					fail = false
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
		if len(verif(Listtostring(showWord), "_")) == 0 {
			for i := 0; i < len(worldbase); i++ {
				print(showWord[i] + " ")
			}
			println("\n\nCongrats !")
			return
		}
	}
	println("You are bad, it's was :", worldbase)
}

func random_word(word string) string {
	file, _ := os.Open(word)
	var word_list []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		word_list = append(word_list, fileScanner.Text())
	}
	return word_list[rand.Intn(len(word))]
}

func verif(word, choice string) []int {
	var ListInd []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[i] {
			ListInd = append(ListInd, i)
		}
	}
	return ListInd
}

func word_choice(mot string) []string {
	var show_word []string
	nbrletter := len(mot)/2 - 1
	for i := 0; i < len(mot); i++ {
		show_word = append(show_word, "_")
	}
	for x := 0; x < nbrletter; x++ {
		ind := rand.Intn(len(mot))
		show_word[ind] = string(mot[ind])
	}
	return show_word
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

func Listtostring(list []string) string {
	char := ""
	for i := 0; i < len(list); i++ {
		char += list[i]
	}
	return char
}

func Lower(choice string) string {
	choice3 := ""
	for h := 0; h < len(choice); h++ {
		if choice[h] >= 65 && choice[h] <= 90 {
			choice3 += string(choice[h] + 32)
		} else {
			choice3 += string(choice[h])
		}
	}
	return choice3
}