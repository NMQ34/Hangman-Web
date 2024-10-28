package hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Définir une structure pour recevoir la donnée de la requête POST
type RequestData struct {
	Text string `json:"text"`
}

// Dictionnaire pour la correspondance entre les lettres accentuées et non accentuées
var accentMap = map[rune][]rune{
	'e': {'e', 'é', 'è', 'ê', 'ë'},
	'a': {'a', 'à', 'â', 'ä'},
	'u': {'u', 'ù', 'û', 'ü'},
	'o': {'o', 'ô', 'ö'},
	'i': {'i', 'î', 'ï'},
	'c': {'c', 'ç'},
}

func Logic(dictionaryPath string) {
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

	dictionaryWords, err := readDictionaryWords(dictionaryPath)
	if err != nil {
		handleError("Erreur lors de la lecture du fichier dictionnaire :", err)
		return
	}

	word, wordRunes := selectRandomWord(dictionaryWords)
	blanks := initializeBlanks(wordRunes)
	guessedLetters := make(map[rune]struct{})
	numLettersToReveal := calculateLettersToReveal(wordRunes)
	revealLetters(wordRunes, blanks, numLettersToReveal)

	lives := 10
	playGame(word, wordRunes, blanks, guessedLetters, hangmanStages, lives, loseMessage, victoryMessage)
}

func handleError(message string, err error) {
	fmt.Println(message, err)
}

func selectRandomWord(dictionaryWords []string) (string, []rune) {
	rand.Seed(time.Now().UnixNano())
	word := dictionaryWords[rand.Intn(len(dictionaryWords))]
	return word, []rune(word)
}

func initializeBlanks(wordRunes []rune) []rune {
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_'
	}
	return blanks
}

func calculateLettersToReveal(wordRunes []rune) int {
	numLettersToReveal := (len(wordRunes) / 2) - 1
	if numLettersToReveal < 0 {
		numLettersToReveal = 0
	}
	return numLettersToReveal
}

func revealLetters(wordRunes []rune, blanks []rune, numLettersToReveal int) {
	revealedIndices := make(map[int]struct{})
	for len(revealedIndices) < numLettersToReveal {
		index := rand.Intn(len(wordRunes))
		revealedIndices[index] = struct{}{}
	}
	for index := range revealedIndices {
		blanks[index] = wordRunes[index]
	}
}

func playGame(word string, wordRunes []rune, blanks []rune, guessedLetters map[rune]struct{}, hangmanStages []string, lives int, loseMessage string, victoryMessage string) {
	for {
		displayGameState(blanks, guessedLetters, lives)
		input := getUserInput()

		if handleInput(input, word, wordRunes, blanks, guessedLetters, &lives) {
			break
		}

		if lives <= 0 {
			fmt.Println(loseMessage)
			fmt.Printf("Perdu! Le mot était : %s\n", word)
			break
		}

		if string(wordRunes) == string(blanks) {
			fmt.Println(victoryMessage)
			fmt.Printf("Le mot était : %s\n", word)
			break
		}
	}
}

func displayGameState(blanks []rune, guessedLetters map[rune]struct{}, lives int) {
	fmt.Printf("Word: [%s], Lettres déjà proposées: %s\n", string(blanks), getGuessedLetters(guessedLetters))
	fmt.Printf("%d essais restants.\n", lives)
}

func getUserInput() string {
	fmt.Print("Entrez une lettre: ")
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(removeAccents(input))
}

func handleInput(input, word string, wordRunes, blanks []rune, guessedLetters map[rune]struct{}, lives *int) bool {
	if len(input) == 0 {
		*lives -= 3
		fmt.Println("Sérieusement !?")
	} else if len(input) > 1 {
		if removeAccents(input) == removeAccents(word) {
			fmt.Println("Vous avez gagné !")
			return true
		} else {
			*lives -= 2
			fmt.Println("Mot incorrect!")
		}
	} else {
		letter := rune(input[0])
		if _, exists := guessedLetters[letter]; exists {
			fmt.Println("Cette lettre a déjà été proposée.")
			return false
		}
		guessedLetters[letter] = struct{}{}
		correctGuess := checkGuess(letter, wordRunes, blanks)
		if !correctGuess {
			*lives--
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

// Fonction pour lire les mots du dictionnaire depuis un fichier
func readDictionaryWords(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Diviser le contenu du fichier en mots, chaque mot étant sur une nouvelle ligne
	words := strings.Split(string(content), "\n")
	return words, nil
}

// Fonction pour retirer les accents des lettres dans une chaîne
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

// Fonction pour afficher les lettres déjà devinées sous forme de chaîne
func getGuessedLetters(m map[rune]struct{}) string {
	letters := make([]string, 0, len(m))
	for k := range m {
		letters = append(letters, string(k))
	}
	return strings.Join(letters, ", ")
}

// Fonction pour vérifier si une lettre non accentuée correspond à une lettre accentuée
func containsAccentMatch(input, wordLetter rune) bool {
	// Vérifier si la lettre sans accent existe dans le dictionnaire des accents
	if accentedLetters, exists := accentMap[input]; exists {
		// Vérifier si la lettre du mot fait partie des lettres accentuées correspondantes
		for _, accentedLetter := range accentedLetters {
			if wordLetter == accentedLetter {
				return true
			}
		}
	}
	// Comparer directement si ce n'est pas une lettre accentuée
	return input == wordLetter
}
