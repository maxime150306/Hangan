package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"golang.org/x/text/unicode/norm"
	"unicode"
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

func enleverAccents(s string) string {
	t := norm.NFD.String(s)
	var result strings.Builder
	for _, r := range t {
		if unicode.Is(unicode.Latin, r) {
			result.WriteRune(r)
		}
	}
	return result.String()
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

func contient(liste []string, c string) bool {
	for _, char := range liste {
		if c == char {
			return true
		}
	}
	return false
}

func Game() {
	mots := chargermots()
	s := strings.TrimSpace(strings.ToLower(enleverAccents(choisirMotaleatoir(mots))))
	motcache := Creermotcache(s)
	motcache = revelerlettresaleatoires(motcache, s, len(s)/2-1)
	var lettrepropose []string

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
		var input string
		fmt.Scanln(&input)
		data.Lettre = strings.TrimSpace(strings.ToLower(enleverAccents(input)))

		if data.Lettre == "" {
			continue
		}
		if len(data.Lettre) != 1 {
			if data.Lettre == s {
				fmt.Println("Bravo vous avez deviné le mot")
				break
			} else {
				try--
				fmt.Println("Mauvaise lettre,il vous reste ", try, " essais")
			}
		}
		if contient(lettrepropose, data.Lettre) {
			fmt.Println("veuillez entrer une lettre qui na pas déjà été proposé")
			continue
		}
		lettrepropose = append(lettrepropose, data.Lettre)

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
				for i := 0; i < 8 && lignesAffichees < len(lignes); i++ {
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
			motCache = motCache[:indice] + enleverAccents(string(mots[indice])) + motCache[indice+1:]
		}
	}
	return motCache
}
