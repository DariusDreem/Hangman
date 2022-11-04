package main

import (
	"bufio"
	"encoding/json"
	ascii_art "ex/ascii-art"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	WordBase   string
	ShowWord   []string
	LetterFind string
	Attemps    int
	Position   int
	Asci       string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//------------------------ launch of program --------------------------------//
	var varia HangManData
	if len(verif(os.Args[1], "--startWith")) > 0 {
		content, _ := os.ReadFile(os.Args[2])
		err1 := json.Unmarshal(content, &varia)
		if err1 != nil {
			print(err1)
			return
		}
	} else {
		word := randomWord(os.Args[1])
		varia = HangManData{word, wordChoice(word), "", 10, -1, ""}
	}
	if len(os.Args[1:]) > 1 && os.Args[2] == "--letterFile" {
		varia.Asci = os.Args[3]
	}
	var letter rune
	var choice string
	fail := false

	println("Good Luck, you have ", varia.Attemps, " attempts.")

	// ------------------------ body of program --------------------------------//
	for varia.Attemps > 0 {
		Affichage(varia.Asci, varia.ShowWord)
		print("\n" + "\n" + "Choose: ")
		fmt.Scanln(&choice)
		choice = Lower(choice)

		if len(choice) == 1 {
			letter = rune(choice[0])
			if len(verif(varia.LetterFind, choice)) > 0 {
				varia.Attemps--
				varia.Position = gallows(1, varia.Position)
				println("\nalready present in the word,", varia.Attemps, "attempts remaining\n")
				fail = true
			}
			varia.LetterFind += choice
			if len(verif(varia.WordBase, choice)) >= 1 {
				index := verif(varia.WordBase, choice)
				for i := 0; i < len(index); i++ {
					varia.ShowWord[index[i]] = string(letter - 32)
				}
			} else {
				if !fail {
					varia.Attemps--
					varia.Position = gallows(1, varia.Position)
					println("\nNot present in the word,", varia.Attemps, "attempts remaining\n")
					fail = false
				}
			}
		} else {
			if choice == varia.WordBase {
				println("\nCongrats !")
				return
			} else if choice == "stop" {
				b, _ := json.Marshal(varia)
				save, _ := os.Create("save.txt")
				save.Write(b)
				return
			} else {
				varia.Attemps -= 2
				varia.Position = gallows(2, varia.Position)
				println("\nlie! is not the real word,", varia.Attemps, "attempts remaining\n")
			}
		}
		if len(verif(ListToString(varia.ShowWord), "_")) == 0 {
			Affichage(varia.Asci, varia.ShowWord)
			println("\n\nCongrats !")
			return
		}
	}
	println("You are bad, it's was :", varia.WordBase)
}

func randomWord(word string) string {
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

func wordChoice(mot string) []string {
	var show_word []string
	nbrletter := len(mot)/2 - 1
	for i := 0; i < len(mot); i++ {
		show_word = append(show_word, "_")
	}
	for x := 0; x < nbrletter; x++ {
		ind := rand.Intn(len(mot))
		show_word[ind] = string(mot[ind] - 32)
	}
	return show_word
}

func gallows(nbr, position int) int {
	jose, _ := os.ReadFile("hangman.txt")
	position += 79 * nbr
	if position >= 710 {
		position = 709
	}
	fmt.Print(string(jose[position-78 : position]))
	return position
}

func ListToString(list []string) string {
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

func Affichage(maj string, word []string) {
	if maj != "" {
		ascii_art.Aff(ListToString(word), maj)
	} else {
		for i := 0; i < len(word); i++ {
			print(word[i] + " ")
		}
	}
}
