package hangmanweb

import (
    "encoding/json"
    "net/http"
    hg "hangmanweb/hangman-classic/functions"
)

type RequestData struct {
    Text string `json:"text"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        http.ServeFile(w, r, "index.html")
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Seule la méthode POST est autorisée", http.StatusMethodNotAllowed)
        return
    }

    var data RequestData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Erreur lors de la lecture de la requête", http.StatusBadRequest)
        return
    }

    // Initialiser et démarrer le jeu
    game := hg.NewGame("Word Selection/dictionnaire.txt")
    game.Start(data.Text, w)

    // Créer la réponse JSON
    response := map[string]interface{}{
        "word":           string(game.Blanks),
        "lives":          game.Lives,
        "guessedLetters": hg.GetGuessedLetters(game.GuessedLetters),
        "message":        "",
    }

    if game.Lives <= 0 {
        response["message"] = "Perdu! Le mot était : " + game.WordToGuess
    } else if string(game.WordToGuess) == string(game.Blanks) {
        response["message"] = "Vous avez gagné !";
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
