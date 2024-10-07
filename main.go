package main

import (
	"fmt"
	hg "hangman/fonction"
	"os"
)

func main() {
	hg.Game("essaie")
	Data,_ := os.ReadFile("hangman.txt")
	fmt.Println(string(Data))
}
