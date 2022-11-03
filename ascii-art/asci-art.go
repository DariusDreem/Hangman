package ascii_art

import (
	"bufio"
	"fmt"
	"os"
)

func Aff() {

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
	fmt.Println(liste)
	file.Close()
}
