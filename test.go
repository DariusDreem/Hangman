package main

func verifint(liste []int, nbr int) bool {
	for i := 0; i < len(liste); i++ {
		if liste[i] == nbr {
			return false
		}
	}
	return true
}

func verif2(word, choice string) []int {
	var listeInd []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[i] {
			listeInd = append(listeInd, i)
		}
	}
	return listeInd
}
