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
	fichier, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return nil // Retourner nil en cas d'erreur
	}
	defer fichier.Close()
	var mots []string

	scanner := bufio.NewScanner(fichier)
	//scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		mots = append(mots, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return mots

}
func chargerpendu(nomFichier string) ([]string, error) {// Chargerpendu est la fonction qui permet de charger le jeu du pendu."Hangman.txt" {
	file, err := os.Open(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier %s: %w", nomFichier, err)// Retourner nil en cas d'erreur
	}
	defer file.Close()
	
	var lignes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lignes = append(lignes, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return lignes, nil
	//scanner.Split(bufio.ScanWords)

	
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
	motcache = revelerlettresaleatoires(motcache, s, 2)

	var data Hangman
	try := 10
	// Lire le fichier de lignes
	lignes, err := chargerpendu("hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier hangman.txt :", err)
		return
	}

	lignesAffichees := 0 // Compteur pour les lignes affichées

	for try > 0 {
		fmt.Println(motcache)
		fmt.Print("Entrez une lettre : ")
		fmt.Scanln(&data.Lettre)

		//vérification de l'entré de l'utilisateur
		data.Lettre = strings.TrimSpace(strings.ToLower(data.Lettre))
		if len(data.Lettre) != 1 {
			fmt.Println("Veuillez entrer une seule lettre")
			continue
		}
		if contient(motcache, data.Lettre) {
			fmt.Println("veuillez entrer une lettre qui na pas encore ete devinée")
			continue
		}

		data.BonneLettre = false
		data.Nvmot = ""

		for i, char := range s {
			if string(char) == data.Lettre {
				data.Nvmot += data.Lettre
				data.BonneLettre = true
			} else {
				data.Nvmot += string(motcache[i])

			}
		}
		motcache = data.Nvmot

		if !data.BonneLettre {
			try--
			fmt.Println("Mauvaise lettre,il vous reste ", try, " essais")
			if lignesAffichees < len(lignes) {
				for i := 0; i < 7 && lignesAffichees < len(lignes); i++ {
					fmt.Println(lignes[lignesAffichees])
					lignesAffichees++
				}
			}
		}
		if motcache == s {
			fmt.Println("Bravo vous avez deviné le mot")
			break
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

func Creermotcache(s string) string {
	motcache := strings.Repeat("_", len(s)) // Crée une chaîne de underscores de la même longueur que le mot
	return motcache
}
func revelerlettresaleatoires(motCache string, mots string, nombredeLettres int) string {
	indicesReverses := make(map[int]bool)

	for len(indicesReverses) < nombredeLettres {
		indice := rand.Intn(len(mots))
		if !indicesReverses[indice] {
			indicesReverses[indice] = true
			motCache = motCache[:indice] + string(mots[indice]) + motCache[indice+1:]
		}
	}
	return motCache
}
/*func parserFichier(nomFichier string) ([]string, error) {
	file, err := os.Open("Hangman.txt")
	if err != nil {
		return nil,
			fmt.Errorf("erreur lors de l'ouverture du fichier %s: %w", nomFichier, err)
	}
	defer file.Close()

	var lignes []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lignes = append(lignes, strings.TrimSpace(scanner.Text())) // Ajouter chaque ligne à la slice
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier %s: %w", nomFichier, err)
	}

	return lignes, nil
}*/
