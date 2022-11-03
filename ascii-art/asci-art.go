package ascii_art

import (
	"bufio"
	"os"
)

func aff() {

	var liste [][]string
	var lettre []string
	nbr := 1

	file, _ := os.Open("standard.txt")

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		lettre = append(lettre, fileScanner.Text())
		if nbr%10 == 0 {
			liste = append(liste, lettre)
			lettre = lettre[len(lettre):]
			nbr = 1
		}
		nbr++
	}

	//print message
	message := "va je ne te hais point"
	for i := 0; i < 9; i++ {

		for z := 0; z < len(message); z++ {
			if message[z]-32 == 0 {
				print("      ")
			} else {
				print(liste[message[z]-32][i])
			}
		}
		print("\n")
	}
	file.Close()
}
