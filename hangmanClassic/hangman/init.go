package hangman

import (
	"fmt"
)

func (g *Game) Init() {

	// Initialisation des variables de l'engine
	g.IsRunning = true

}

// Fonction pour créer une nouvelle instance de jeu
func (g *Game) NewGame(dictionaryPath string) {
	dictionaryWords, err := readDictionaryWords(dictionaryPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dictionnaire :", err)
	}

	word, _ := selectRandomWord(dictionaryWords)
	return g.Player{
		WordToGuess:    word,
		Blanks:         g.initializeBlanks([]rune(word)),
		Lives:          10,
		GuessedLetters: make(map[rune]struct{}),
		IsAlive:        true,
	}
}

func (g *Game) InitPlayer() {

	g.Player = Player{
		WordToGuess:    word,
		Blanks:         initializeBlanks([]rune(word)),
		Lives:          10,
		GuessedLetters: make(map[rune]struct{}),
		IsAlive:        true,
	}
}
