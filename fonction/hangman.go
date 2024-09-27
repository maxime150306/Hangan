package hangman

import "fmt"

func Hangman(s string) {
	var try int = 10
	var motcache string
	for i := 0; i < len(s); i++ {
		motcache = motcache + "_"
	}
	for try > 0 {
		fmt.Println(motcache)
		var lettre string
		fmt.Print("Entrez une lettre : ")
		_, err := fmt.Scanf("%s", &lettre)
		if err != nil {

		}
		var nvmot string
		for i, char := range s {
			if string(char) == lettre {
				nvmot = nvmot + lettre
			} else {
				nvmot = nvmot + string(motcache[i])
			}
		}
		motcache = nvmot
		fmt.Println(motcache)
		
	}
}
