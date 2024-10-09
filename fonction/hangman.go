package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Game est la fonction qui permet de lancer le jeu du pendu.
// elle prend en paramètre un string qui est le mot à trouver.
// elle lance le jeu et gère les interactions avec l'utilisateur.
// elle affiche le mot à trouver avec des _ à la place des lettres à trouver.
// elle demande à l'utilisateur d'entrer une lettre.
// si la lettre est bonne, elle remplace le _ par la lettre.
// si la lettre est mauvaise, elle enlève un essai.
// si l'utilisateur a plus d'essai, elle affiche un message de défaite.
// sinon, elle affiche un message de victoire.
func timeNow() time.Time {
	return time.Now()
}
func chargermots() []string {
	fichier, _ := os.Open("words.txt")
	defer fichier.Close()
	var mots []string

	scanner := bufio.NewScanner(fichier)
	scanner.Split(bufio.ScanWords)
	mots = append(mots, strings.TrimSpace(scanner.Text()))
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}
	return mots

}

func choisirMotaleatoir(mots []string) string {
	rand.Seed(timeNow().UnixNano())
	return mots[rand.Intn(len(mots))]
}

func contient(s string, c string) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == c {
			return true
		}
	}
	return false
}
func Game() {
	mots := chargermots()
	s := strings.TrimSpace(strings.ToLower(choisirMotaleatoir(mots)))
	motcache := Creermotcache(s)
	motcache = revelerlettresaleatoires(motcache,s, 2)
	if s == "" {
		fmt.Println("Veuillez entrer un mot")
		return
	}
	var data Hangman
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
		data.Lettre = strings.TrimSpace(strings.ToLower(data.Lettre))
		if !data.BonneLettre {
			try--
			fmt.Println("Mauvaise lettre,il vous reste ", try, " essais")
		}
		if motcache == s {
			fmt.Println("Bravo")
			break
		}
		if contient(motcache, data.Lettre) {
			fmt.Println("veuillez entrer une lettre qui na pas encore ete devinée")
		}
		if try == 0 {
			fmt.Println("Perdu le mot est :", s)
			break
		}
		if data.Lettre == motcache {
			fmt.Println("Bravo")
			break
		}
	}
}
func revelerlettresaleatoires(motCache string, mots string, nombredeLettres int) string {
	indicesReverses := make(map[int]bool)
	motCache = strings.Repeat("_", len(mots))
	for len(indicesReverses) < nombredeLettres {
		indice := rand.Intn(len(mots))
		if !indicesReverses[indice] {
			indicesReverses[indice] = true
			motCache = motCache[:indice] + string(mots[indice]) + motCache[indice+1:]
		}
	}
	return motCache
}
func Creermotcache(s string) string {
	var data Hangman
	data.Motcache = ""
	for i := 0; i < len(data.Motcache); i++ {
		data.Motcache = data.Motcache + "_"
	}
	return data.Motcache
}
