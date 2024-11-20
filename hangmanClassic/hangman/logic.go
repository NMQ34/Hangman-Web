package hangman

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func (s *Structure) Verif(letter rune) bool {
	isHere := false
	for _, i := range s.LetterTested {
		if i == string(letter) {
			isHere = true
			return false
		}
	}
	if !isHere {
		s.LetterTested = append(s.LetterTested, string(letter))
	}
	return true
}

func (s *Structure) VWord(letter rune) {
	isHere := false
	for index, r := range s.SecretWord {
		if rune(letter) == r {
			s.Blanks[index] = s.Blanks[letter]
			isHere = true
		}
	}
	if !isHere {
		s.Lives -= 1
	}
}

// Sélectionner un mot au hasard
func (s *Structure) SelectRandomWord() []rune {
	rand.Seed(time.Now().UnixNano())
	content, err := ioutil.ReadFile("hangmanClassic/texte/dictionnaire.txt")
	if err != nil {
		return nil
	}
	ListOfWord := strings.Split(string(content), "\n")
	randomIndex := rand.Intn(len(ListOfWord))
	word := ListOfWord[randomIndex]
	return []rune(word)
}

// Initialiser les blancs
func (s *Structure) InitializeBlanks() []rune {
	blanks := make([]rune, len(s.SecretWord))
	for i := range blanks {
		blanks[i] = '_'
	}
	return blanks
}

func (s *Structure) CheckOut() {
	if s.Lives == 0 {
		s.Running = false
		s.Lose = true
	}
	var count = len(s.SecretWord)
	var count2 int = 0
	for _, i := range s.SecretWord {
		if s.Blanks[i] == rune(i) {
			count2 += 1
		}
	}
	if count2 == count {
		s.Win = true
	}
}
