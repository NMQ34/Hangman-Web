package hangman

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// VerifLetter vérifie si la lettre proposée a déjà été testée.
// Si elle n'a pas été testée, elle est ajoutée à `LetterTested`. La fonction retourne `true`
// si la lettre est nouvelle, et `false` sinon.
func (s *Structure) VerifLetter(letter []rune) bool {
	isHere := true // Indique que la lettre n'est pas encore dans LetterTested.

	// Parcourt les lettres déjà testées pour vérifier si la lettre est déjà présente.
	for _, i := range s.LetterTested {
		if i == letter[0] {
			isHere = false // La lettre est déjà présente.
		}
	}

	// Si la lettre est nouvelle, elle est ajoutée à LetterTested.
	if isHere && len(s.LetterTested) == 0 {
		s.LetterTested += string(letter)
	} else if isHere {
		s.LetterTested += ", " + string(letter)
	}

	return isHere
}

// CheckLetter vérifie si une lettre est présente dans le mot secret.
// Si la lettre est trouvée, elle est insérée dans `Blanks` à la bonne position.
// La fonction retourne `true` si la lettre est dans le mot, sinon `false`.
func (s *Structure) CheckLetter(letter []rune) bool {
	isHere := false

	// Parcourt les lettres du mot secret pour vérifier une correspondance.
	for _, l := range letter {
		for index, r := range s.SecretWord {
			if l == r {
				s.Blanks[index] = r // Remplace le tiret par la lettre trouvée.
				isHere = true
			}
		}
	}

	// Vérifie si la lettre a déjà été testée.
	if !s.VerifLetter(letter) {
		isHere = true
	}

	return isHere
}

// ConvertRinS convertit un tableau de runes en une chaîne de caractères.
func (s *Structure) ConvertRinS(runes []rune) string {
	var chaine string = ""

	// Ajoute chaque rune à une chaîne.
	for _, i := range runes {
		chaine += string(i)
	}

	return chaine
}

// CheckWord vérifie si le mot complet proposé par le joueur est correct.
// Si toutes les lettres correspondent au mot secret, la partie est gagnée.
func (s *Structure) CheckWord(letter []rune) bool {
	var isHere bool = true

	// Vérifie chaque lettre du mot proposé.
	for _, l := range letter {
		if !s.CheckLetter([]rune{l}) {
			isHere = false
		}
	}

	// Si toutes les lettres correspondent, marque la victoire.
	if isHere {
		s.Win = true
	}

	return isHere
}

// SelectRandomWord sélectionne un mot au hasard depuis un fichier texte.
// Le mot est retourné sous forme de tableau de runes.
func (s *Structure) SelectRandomWord() []rune {
	// Initialise une graine pour le générateur aléatoire.
	rand.Seed(time.Now().UnixNano())

	// Lit le contenu du fichier contenant la liste des mots.
	content, err := os.ReadFile("texte/dictionnaire.txt")
	if err != nil {
		return nil // Retourne nil en cas d'erreur de lecture.
	}

	// Divise le contenu du fichier en une liste de mots.
	ListOfWord := strings.Split(string(content), "\n")
	randomIndex := rand.Intn(len(ListOfWord)) // Sélectionne un mot au hasard.
	str := ListOfWord[randomIndex]
	word := str[:len(str)-1] // Retire le caractère de fin de ligne.

	return []rune(word)
}

// InitializeBlanks crée un tableau de tirets bas "_" correspondant à la longueur du mot secret.
// Chaque position représente une lettre du mot secret.
func (s *Structure) InitializeBlanks() []rune {
	blanks := make([]rune, len(s.SecretWord))

	// Remplit le tableau avec des tirets bas.
	for i := range blanks {
		blanks[i] = '_'
	}

	return blanks
}

// CheckOut vérifie si le joueur a perdu (plus de vies) ou gagné (mot complété).
func (s *Structure) CheckOut() {
	// Si le joueur n'a plus de vies, la partie est terminée avec une défaite.
	if s.Lives == 0 {
		s.Running = false
		s.Lose = true
	}

	// Compte les lettres correctement trouvées dans Blanks.
	var count = len(s.SecretWord)
	var count2 int = 0

	for index, i := range s.SecretWord {
		if s.Blanks[index] == rune(i) {
			count2 += 1
		}
	}

	// Si toutes les lettres sont découvertes, la partie est gagnée.
	if count2 == count {
		s.Win = true
	}
}
