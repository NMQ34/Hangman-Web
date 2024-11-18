package hangman

import (
	"encoding/json"
	"net/http"
)

// Structure pour le jeu du pendu
type Player struct {
	WordToGuess    string
	Blanks         []rune
	Lives          int
	GuessedLetters map[rune]struct{}
	IsAlive        bool
}

// Structure pour recevoir les données de la requête POST
type RequestData struct {
	Text string `json:"text"`
}

type Game struct {
	Player    Player
	IsRunning bool
}

// Handler pour gérer la racine "/"
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Gérer la requête POST pour le jeu
		var data RequestData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture de la requête", http.StatusBadRequest)
			return
		}

	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
