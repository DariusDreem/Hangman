package main

import "fmt"

func main() {
	var liste []int
	liste = append(liste, 6)
	liste = append(liste, 2)
	liste = append(liste, 5)
	liste = append(liste, 7)
	nbr := 5
	val := search_nbr(liste, nbr)
	fmt.Print(val)
}
func search_nbr(param []int, nbr int) bool {
	for i := 0; i < len(param); i++ {
		if param[i] == nbr {
			return false
		}
	}
	return true
}
