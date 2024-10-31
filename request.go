package main

import (
	hangman "hangmanweb/hangman-classic/functions"
	"html/template"
	"net/http"
)

var hc *hangman.Game

/*WordToGuess    string
Blanks         []rune
Lives          int
GuessedLetters map[rune]struct{}*/

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
	Alphabet()
	Data.BoolValue = false
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/input", Input)
	http.ListenAndServe(":8080", nil)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, hc)
}

func Input(w http.ResponseWriter, r *http.Request) {
	hc.WordToGuess = r.FormValue("textearecup")
	Data.BoolValue = true

	MainPage(w, r)
}

func Alphabet() {
	for i := 'a'; i != 'z'+1; i++ {
		Data.NumberOfButton = append(Data.NumberOfButton, Range{Letter: string(i)})
	}
}

/*package hangmanweb

import (
    "encoding/json"
    "net/http"
    "html/template"
    hg "hangmanweb/hangman-classic/functions"
)

type RequestData struct {
    Text string `json:"text"`
}
type donnee struct {
    Lives int

}
var d = donnee{
    Lives : game.Lives,
}
func Handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, d)
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
*/
