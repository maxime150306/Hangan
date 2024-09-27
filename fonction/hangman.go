package hangman

import "fmt"

// Hangman est une fonction qui prend une string en parametre et qui permet au joueur 
// de trouver le mot en devinant des lettres. Si le joueur se trompe, il perd une vie.
// Lorsque le joueur a trouve le mot, il a gagne. Lorsqu'il a perdu toutes ses vies, 
// il a perdu.
func Hangman(s string) {
	var try int = 10
	motcache := Creermotcache(s)
	for try > 0 {
		fmt.Println(motcache)
		var lettre string
		fmt.Print("Entrez une lettre : ")
		fmt.Scanf("%s", &lettre)

		var nvmot string
		var bonneLettre bool = false
		for i, char := range s {
			if string(char) == lettre {
				nvmot = nvmot + lettre
				bonneLettre = true
			} else {
				nvmot = nvmot + string(motcache[i])
			}

		}
		motcache = nvmot
		if !bonneLettre {
			try = try - 1
			fmt.Println("Mauvaise lettre, il vous reste", try, "vies")
		}	
		if motcache == s {
			fmt.Println("Bravo, vous avez gagné !")
			break
		}
		if try == 0 {
			fmt.Println("Dommage, vous avez perdu !")
			break
		}
		
	}
}

func Creermotcache(s string) string {
	var motcache string
	for i := 0; i < len(s); i++ {
		motcache = motcache + "_"
	}
	return motcache
}