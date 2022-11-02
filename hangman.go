package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	jo, _ := ioutil.ReadFile("hangman.txt")
	position := 0
	var choix string
	var letter rune
	LetterFind := ""
	mot := motrandom(os.Args[1])
	attemps := 10
	println(mot)
	println(len(mot))
	println("Good Luck, you have ", attemps, " attempts.")
	mot_cache := creadumot(mot)
	echec := false
	for attemps > 0 {
		for i := 0; i < len(mot); i++ {
			print(mot_cache[i] + " ")
		}
		print("\n" + "\n" + "Choose: ")
		fmt.Scanln(&choix)
		var listeind []int
		for i := 0; i < len(mot); i++ {
			if choix[0] == mot[i] {
				listeind = append(listeind, i)
			}
		}
		if len(choix) == 1 {
			letter = rune(choix[0])
			if len(verif(LetterFind, choix)) > 0 {
				attemps--
				position = pendu(1, position)
				println("\nalready present in the word,", attemps, "attempts remaining\n")
				echec = true
			}
			LetterFind += choix
			if len(verif(worldbase, choix)) >= 1 {
				for i := 0; i < len(index); i++ {
					showWord[index[i]] = string(letter - 32)
				}
			} else {
				if !echec {
					attemps--
					position = pendu(1, position)
					println("\nNot present in the word,", attempsttemps, "attempts remaining\n")
					echec = false
		}else {
					if choix == worldBase {
						println("\nCongrats !")
						return
					} else if choix == "stop" {
						b, _ := json.Marshal(varia)
						save, _ := os.Create("save.txt")
						save.Write(b)
						return
					}
		}
		if len(listeind) > 0 {
			for k := 0; k < len(listeind); k++ {
				mot_cache[listeind[k]] = choix
			}
		} else {
			attemps--
			println("\nNot present in the word,", attemps, "attempts remaining\n")
			fmt.Println(string(jo[position : position+80]))
			position += 79
		}
	}
	println("t'es nul c'était :", mot)
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
