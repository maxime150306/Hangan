package main

import "fmt"

func Hangman(s string)  {
	
}

func VerifLettre(lettre string, s string, motcache[]string) int {
	for i, char := range s {
		if string(char) == lettre {
			motcache[i] = lettre
		}
	}
	return -1
}

func main() {
	var mot string = "essaie"
	var motcache []string
	var motl []string
	for i := 0; i < len(mot); i++ {
		motcache = append(motcache, "_")
	}
	for _, char := range mot {
		motl = append(motl, string(char))
	}
	lettre := "e"
	fmt.Println(motl)
	fmt.Println(motcache)
	VerifLettre(lettre, mot, motcache)
	fmt.Println(motcache)

}
