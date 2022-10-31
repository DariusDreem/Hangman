package main

func verifie(liste []int, nbr int) bool {
	for i := 0; i < len(liste); i++ {
		if liste[i] == nbr {
			return false
		}
	}
	return true
}
