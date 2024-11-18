package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Range struct {
	Letter string
}

type DataForm struct {
	Word           string
	BoolValue      bool
	NumberOfButton []Range
}

var Data DataForm

func main() {
	var hc *hangman.Game
	hc.Init()
	// Initialisation du jeu ici avant de démarrer le serveur
	hc = hc.Game.NewGame("Word Selection/dictionnaire_hardmode.txt")

	// Initialisation des composants du jeu
	game := Game.NewGame("hangmanClassic/texte/dictionnaire.txt")
	game.Start(data.Text, w)
	if hc == nil {
		fmt.Println("Erreur lors de l'initialisation du jeu.")
		return
	}
	Alphabet() // Initialiser l'alphabet pour l'affichage des boutons

	// Initialiser les routes du serveur
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/input", Input)

	// Démarrer le serveur HTTP
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}

// Page d'accueil du jeu
func MainPage(w http.ResponseWriter, r *http.Request) {
	// Utiliser le template HTML pour afficher le jeu
	tmpl := template.Must(template.ParseFiles("menu.html"))
	tmpl.Execute(w, hc)
}

// Fonction pour gérer les entrées de l'utilisateur
func Input(w http.ResponseWriter, r *http.Request) {
	letter := r.FormValue("textearecup") // Récupérer la lettre entrée

	// Gérer la lettre devinée
	win = hc.Game.HandleInput(letter)

	// Préparer la réponse JSON
	response := map[string]interface{}{
		"word":           string(hc.Blanks),
		"lives":          hc.Lives,
		"guessedLetters": hangman.Game.GetGuessedLetters(Player.GuessedLetters),
		"message":        "",
	}

	// Vérifier si le joueur a perdu ou gagné
	if hangmanClassic.Player.Lives <= 0 {
		response["message"] = "Perdu! Le mot était : " + hc.WordToGuess
	} else if win {
		response["message"] = "Vous avez gagné !"
	}

	// Renvoyer la réponse en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Fonction pour initialiser l'alphabet
func Alphabet() {
	for i := 'a'; i <= 'z'; i++ {
		Data.NumberOfButton = append(Data.NumberOfButton, Range{Letter: string(i)})
	}
}
