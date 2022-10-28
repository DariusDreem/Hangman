package main

import (
	"fmt"
)

func main() {
	var choix string
	mot := "teste"
	var mot_cache []string
	attemps := 10
	var indice int
	println("Good Luck, you have ", attemps, " attempts.")
	for attemps > 0 {
		for i := 0; i < len(mot); i++ {
			mot_cache = append(mot_cache, "_")
		}
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
		//test j'espere que c bon
		indice = -1
	}
}
