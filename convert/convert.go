package convert

func Listtostring(list []string) string {
	chaine := ""
	for i := 0; i < len(list); i++ {
		chaine += list[i]
	}
	return chaine
}

func minuscule(choice string) string {
	choice3 := ""
	for h := 0; h < len(choice); h++ {
		if choice[h] >= 65 && choice[h] <= 90 {
			choice3 += string(choice[h] + 32)
		} else {
			choice3 += string(choice[h])
		}
	}
	return choice3
}
