package hangman

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structure pour le jeu du pendu
type Game struct {
	WordToGuess    string
	Blanks         []rune
	Lives          int
	GuessedLetters map[rune]struct{}
}

// Structure pour recevoir les données de la requête POST
type RequestData struct {
	Text string `json:"text"`
}

// Handler pour gérer la racine "/"
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Afficher une simple page HTML en réponse à une requête GET
		fmt.Fprintf(w, `<html>
			<head><title>Hangman Game</title></head>
			<body>
				<h1>Bienvenue au jeu du Pendu !</h1>
				<p>Utilisez une requête POST pour envoyer des lettres ou deviner le mot.</p>
			</body>
			</html>`)
		return
	} else if r.Method == http.MethodPost {
		// Gérer la requête POST pour le jeu
		var data RequestData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture de la requête", http.StatusBadRequest)
			return
		}

		// Initialiser le jeu avec le mot du dictionnaire
		game := NewGame("dictionnaire.txt")
		game.Start(data.Text, w)
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

// Fonction pour créer une nouvelle instance de jeu
func NewGame(dictionaryPath string) *Game {
	dictionaryWords, err := readDictionaryWords(dictionaryPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dictionnaire :", err)
		return nil
	}

	word, _ := selectRandomWord(dictionaryWords)
	return &Game{
		WordToGuess:    word,
		Blanks:         initializeBlanks([]rune(word)),
		Lives:          10,
		GuessedLetters: make(map[rune]struct{}),
	}
}
