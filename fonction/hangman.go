package hangman

import (
	"fmt"
)

func Game(s string) {
	if s == "" {
		fmt.Println("Veuillez entrer un mot")
		return
	}

	var data Hangman
	motcache := Creermotcache(s)
	try := 10

	for try > 0 {
		fmt.Println(motcache)
		fmt.Print("Entrez une lettre : ")
		fmt.Scanln(&data.Lettre)
		//vérification de l'entré de l'utilisateur
		if len(data.Lettre) != 1 {
			fmt.Println("Veuillez entrer une seule lettre")
			continue
		}

		data.BonneLettre = false
		data.Nvmot = ""
		for i, char := range s {
			if string(char) == data.Lettre {
				data.Nvmot = data.Nvmot + data.Lettre
				data.BonneLettre = true
			} else {
				data.Nvmot = data.Nvmot + string(motcache[i])

			}
		}
		motcache = data.Nvmot
		if !data.BonneLettre {
			try--
			fmt.Println("Mauvaise lettre,il vous reste ", try, " essais")
		}
		if motcache == s {
			fmt.Println("Bravo")
			break
		}
		if try == 0 {
			fmt.Println("Perdu")
			break
		}
	}
}

func Creermotcache(s string) string {
	var data Hangman
	data.Motcache = ""
	for i := 0; i < len(s); i++ {
		data.Motcache = data.Motcache + "_"
	}
	return data.Motcache
}
