package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func motaletaoire(file string) string {
	rand.Seed(time.Now().UnixNano())
	content, _ := ioutil.ReadFile(file)
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
	nbr1 := rand.Intn(len(liste))
	return liste[nbr1]
}

func main() {
	fmt.Println(motaletaoire(os.Args[1]))
	jo, _ := ioutil.ReadFile("hangman.txt")
	position := 0
	rand.Seed(time.Now().UnixNano())
	content, _ := ioutil.ReadFile(os.Args[1])
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
	println(mot)
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
			fmt.Println(string(jo[position : position+70]))
			position += 71
		}
	}
	println("t'es nul c'était :", mot)
}
