package main

func verifie(liste []int, nbr int) bool {
	for i := 0; i < len(liste); i++ {
		if liste[i] == nbr {
			return false
		}
	}
	return true
}

func verification(word, choice string) []int {
	var listeInd []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[i] {
			listeInd = append(listeInd, i)
		}
	}
	return listeInd
}
