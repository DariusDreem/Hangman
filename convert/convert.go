package convert

func Listtostring(list []string) string {
	chaine := ""
	for i := 0; i < len(list); i++ {
		chaine += list[i]
	}
	return chaine
}

func minuscule(choix string) string {
	choix3 := ""
	for h := 0; h < len(choix); h++ {
		if choix[h] >= 65 && choix[h] <= 90 {
			choix3 += string(choix[h] + 32)
		} else {
			choice3 += string(choice[h])
		}
	}
	return choix3
}
