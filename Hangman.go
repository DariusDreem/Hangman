package main

import (
	"bufio"
	ascii_art "ex/ascii-art"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	ascii_art.Aff()
	rand.Seed(time.Now().UnixNano())
	jo, _ := ioutil.ReadFile("hangman.txt")
	position := 0
	var choix string
	mot := motrandom(os.Args[1])
	nbrlettre := len(mot)/2 - 1
	var mot_cache []string
	attemps := 10
	println(mot)
	println(len(mot))

	println("Good Luck, you have ", attemps, " attempts.")
	for i := 0; i < len(mot); i++ {
		mot_cache = append(mot_cache, "_")
	}
	for x := 0; x < nbrlettre; x++ {
		ind := rand.Intn(len(mot))
		mot_cache[ind] = string(mot[ind])
	}

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
