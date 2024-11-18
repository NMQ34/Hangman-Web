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
	hangmanStages, err := readHangmanStages("texte/hangman.txt")
	if err != nil {
		g.handleError("Erreur lors de la lecture du fichier :", err)
		return
	}

	g.playGame(hangmanStages)
}

func (g *Game) handleError(message string, err error) {
	fmt.Println(message, err)
}

func (g *Game) playGame(hangmanStages []string) string {
	for {
		g.displayGameState(g.Player.Blanks, g.Player.GuessedLetters, g.Player.Lives)
		input := g.getUserInput()

		if g.handleInput(input) {
			break
		}

		if g.Player.Lives <= 0 {
			fmt.Printf("Perdu! Le mot était : %s\n", g.Player.WordToGuess)
			break
		}

		if string(g.Player.WordToGuess) == string(g.Player.Blanks) {
			fmt.Printf("Le mot était : %s\n", g.Player.WordToGuess)
			break
		}
	}
	return ""
}

func (g *Game) displayGameState(blanks []rune, guessedLetters map[rune]struct{}, lives int) {
	fmt.Printf("Word: [%s], Lettres déjà proposées: %s\n", string(blanks), g.GetGuessedLetters(guessedLetters))
	fmt.Printf("%d essais restants.\n", lives)
}

func (g *Game) getUserInput() string {
	fmt.Print("Entrez une lettre: ")
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(removeAccents(input))
}

func (g *Game) handleInput(input string) bool {
	if len(input) == 0 {
		g.Player.Lives -= 3
		fmt.Println("Sérieusement !?")
	} else if len(input) > 1 {
		if removeAccents(input) == removeAccents(g.Player.WordToGuess) {
			fmt.Println("Vous avez gagné !")
			return true
		} else {
			g.Player.Lives -= 2
			fmt.Println("Mot incorrect!")
		}
	} else {
		letter := rune(input[0])
		if _, exists := g.Player.GuessedLetters[letter]; exists {
			fmt.Println("Cette lettre a déjà été proposée.")
			return false
		}
		g.Player.GuessedLetters[letter] = struct{}{}
		correctGuess := g.checkGuess(letter, []rune(g.Player.WordToGuess), g.Player.Blanks)
		if !correctGuess {
			g.Player.Lives--
			fmt.Println("Lettre incorrecte!")
		}
	}
	return false
}

func (g *Game) checkGuess(letter rune, wordRunes []rune, blanks []rune) bool {
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

// Fonction pour afficher les lettres déjà devinées sous forme de chaîne
func (g *Game) GetGuessedLetters(m map[rune]struct{}) string {
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
func (g *Game) initializeBlanks(wordRunes []rune) []rune {
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
