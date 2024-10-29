package hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time" // Ajout de l'importation pour le package time
)

// Variable pour les accents
var accentMap = map[rune][]rune{
	'e': {'e', 'é', 'è', 'ê', 'ë'},
	'a': {'a', 'à', 'â', 'ä'},
	'u': {'u', 'ù', 'û', 'ü'},
	'o': {'o', 'ô', 'ö'},
	'i': {'i', 'î', 'ï'},
	'c': {'c', 'ç'},
}

// Fonction pour démarrer le jeu
func (g *Game) Start(inputText string, w http.ResponseWriter) {
	// Démarrer la logique de jeu ici
	hangmanStages, err := readHangmanStages("ASCIIDisplay/hangman.txt")
	if err != nil {
		handleError("Erreur lors de la lecture du fichier :", err)
		return
	}

	loseMessage, err := readLoseMessage("ASCIIDisplay/lose.txt")
	if err != nil {
		handleError("Erreur lors de la lecture du fichier de lose :", err)
		return
	}

	victoryMessage, err := readVictoryMessage("ASCIIDisplay/victory.txt")
	if err != nil {
		handleError("Erreur lors de la lecture du fichier de victoire :", err)
		return
	}

	playGame(g, hangmanStages, loseMessage, victoryMessage)
}

func handleError(message string, err error) {
	fmt.Println(message, err)
}

func playGame(game *Game, hangmanStages []string, loseMessage string, victoryMessage string) {
	for {
		displayGameState(game.Blanks, game.GuessedLetters, game.Lives)
		input := getUserInput()

		if handleInput(input, game) {
			break
		}

		if game.Lives <= 0 {
			fmt.Println(loseMessage)
			fmt.Printf("Perdu! Le mot était : %s\n", game.WordToGuess)
			break
		}

		if string(game.WordToGuess) == string(game.Blanks) {
			fmt.Println(victoryMessage)
			fmt.Printf("Le mot était : %s\n", game.WordToGuess)
			break
		}
	}
}

func displayGameState(blanks []rune, guessedLetters map[rune]struct{}, lives int) {
	fmt.Printf("Word: [%s], Lettres déjà proposées: %s\n", string(blanks), GetGuessedLetters(guessedLetters))
	fmt.Printf("%d essais restants.\n", lives)
}

func getUserInput() string {
	fmt.Print("Entrez une lettre: ")
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(removeAccents(input))
}

func handleInput(input string, game *Game) bool {
	if len(input) == 0 {
		game.Lives -= 3
		fmt.Println("Sérieusement !?")
	} else if len(input) > 1 {
		if removeAccents(input) == removeAccents(game.WordToGuess) {
			fmt.Println("Vous avez gagné !")
			return true
		} else {
			game.Lives -= 2
			fmt.Println("Mot incorrect!")
		}
	} else {
		letter := rune(input[0])
		if _, exists := game.GuessedLetters[letter]; exists {
			fmt.Println("Cette lettre a déjà été proposée.")
			return false
		}
		game.GuessedLetters[letter] = struct{}{}
		correctGuess := checkGuess(letter, []rune(game.WordToGuess), game.Blanks)
		if !correctGuess {
			game.Lives--
			fmt.Println("Lettre incorrecte!")
		}
	}
	return false
}

func checkGuess(letter rune, wordRunes []rune, blanks []rune) bool {
	correctGuess := false
	for i, wordLetter := range wordRunes {
		if containsAccentMatch(letter, wordLetter) {
			blanks[i] = wordLetter
			correctGuess = true
			fmt.Println("Bon choix")
		}
	}
	return correctGuess
}

// Fonction pour lire les étapes du pendu depuis un fichier
func readHangmanStages(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Diviser le fichier en étapes, chaque étape est séparée par une ligne vide
	stages := strings.Split(string(content), "\n\n")
	return stages, nil
}

// Fonction pour lire le message de défaite depuis un fichier
func readLoseMessage(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Fonction pour lire le message de victoire depuis un fichier
func readVictoryMessage(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Fonction pour afficher les lettres déjà devinées sous forme de chaîne
func GetGuessedLetters(m map[rune]struct{}) string {
	letters := make([]string, 0, len(m))
	for k := range m {
		letters = append(letters, string(k))
	}
	return strings.Join(letters, ", ")
}

// Lire les mots du dictionnaire depuis un fichier
func readDictionaryWords(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(content), "\n")
	return words, nil
}

// Sélectionner un mot au hasard
func selectRandomWord(dictionaryWords []string) (string, []rune) {
	rand.Seed(time.Now().UnixNano())
	word := dictionaryWords[rand.Intn(len(dictionaryWords))]
	return word, []rune(word)
}

// Initialiser les blancs
func initializeBlanks(wordRunes []rune) []rune {
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_'
	}
	return blanks
}

// Supprimer les accents d'une chaîne
func removeAccents(input string) string {
	replacer := strings.NewReplacer(
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"à", "a", "â", "a", "ä", "a",
		"ù", "u", "û", "u", "ü", "u",
		"ô", "o", "ö", "o",
		"î", "i", "ï", "i",
		"ç", "c",
	)
	return replacer.Replace(input)
}

// Fonction pour vérifier si une lettre non accentuée correspond à une lettre accentuée
func containsAccentMatch(input, wordLetter rune) bool {
	// Vérifier si la lettre sans accent existe dans le dictionnaire des accents
	if accentedLetters, exists := accentMap[input]; exists {
		// Vérifier si la lettre du mot fait partie des lettres accentuées correspondantes
		for _, accentedLetter := range accentedLetters {
			if accentedLetter == wordLetter {
				return true
			}
		}
	}
	return false
}
