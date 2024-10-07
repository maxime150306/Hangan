package hangman

import (
	"bufio"
	"fmt"
	"os"
)

// Hangman est une fonction qui prend une string en parametre et qui permet au joueur
// de trouver le mot en devinant des lettres. Si le joueur se trompe, il perd une vie.
// Lorsque le joueur a trouve le mot, il a gagne. Lorsqu'il a perdu toutes ses vies,
// il a perdu.
func Game(s string) {
	var data Hangman
	motcache := Creermotcache(s)
	Try := 10

	for data.Try > 0 {
		fmt.Println(data.Nvmot)
		fmt.Print("Entrez une lettre : ")
		fmt.Scanf("%s", &data.Lettre)

		for i, char := range s {
			if string(char) == data.Lettre {
				data.Nvmot = data.Nvmot + data.Lettre
				data.BonneLettre = true
			} else {
				data.Nvmot = data.Nvmot + string(motcache[i])
			}
			motcache = data.Nvmot //met a jour la variable motcache
		}

		if !data.BonneLettre {
			data.Try = data.Try - 1
			data.Compteurlignes = data.Compteurlignes + 7

			if data.Compteurlignes > len(data.Lignes) {

				for i := 0; i < len(data.Lignes); i++ {
					fmt.Println(data.Lignes[i])
				}
				fmt.Println("Mauvaise lettre, il vous reste", Try, "vies")

				if data.Try == 0 {
					fmt.Println("Dommage, vous avez perdu !")
					fmt.Println() // affiche la dernière étape du pendu
					break
				}
			}
			for data.Try > 0 {
				file, err := os.Open("hangman.txt")
				if err != nil {
					fmt.Println("Erreur lors de l'ouverture du fichier : ", err)
					return
				}
				defer file.Close()
				fmt.Println(data.Motcache)
				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanLines)
				fmt.Println(scanner)
				Lignes := make([]string, 0)
				data.Compteurlignes = len(Lignes)
				Compteur := 0
				for scanner.Scan() {
					Lignes = append(Lignes, scanner.Text())
				}

				for i := 0; i < data.Compteurlignes; i++ {
					fmt.Println(Lignes[i])
					for j := 0; j < 7 && i+j < Compteur; j++ {
						fmt.Print(Lignes[i+j])

						fmt.Println()
						Compteur++
						break
					}
				}

				if motcache == s {
					fmt.Println("Bravo, vous avez gagné !")
					break

				}

				Data, _ := os.ReadFile("hangman.txt")
				fmt.Println(string(Data))
			}

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
