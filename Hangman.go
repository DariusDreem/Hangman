package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	content, _ := ioutil.ReadFile("words.txt")
	chaine := ""
	var liste []string
	for i := 0; i < len(content); i++ {
		if content[i] == 10 {
			liste = append(liste, chaine)
			chaine = ""
		} else {
			chaine += string(content[i])
		}
	}
	var choix string
	nbr1 := rand.Intn(len(liste))
	mot := liste[nbr1]
	nbrlettre := len(mot)/2 - 1
	println(nbrlettre)
	var mot_cache []string
	attemps := 10
	var indice int
	println(mot)
	println("Good Luck, you have ", attemps, " attempts.")
	for i := 0; i < len(mot); i++ {
		mot_cache = append(mot_cache, "_")
	}

	for attemps > 0 {
		for i := 0; i < len(mot); i++ {
			print(mot_cache[i] + " ")
		}
		print("\n" + "\n" + "Choose: ")
		fmt.Scanln(&choix)
		for i := 0; i < len(mot); i++ {
			if choix[0] == mot[i] {
				indice = i
				break
			}
		}
		if indice != -1 {
			mot_cache[indice] = choix
		} else {
			attemps--
			println("\nNot present in the word,", attemps, "attempts remaining\n")
		}
		indice = -1
	}
	println("t'es nul c'était :", mot)
}
