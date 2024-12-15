package hangman

// init initialise les variables nécessaires pour démarrer une nouvelle partie du jeu du Pendu.
// Elle configure les états initiaux comme le mot à deviner, les vies disponibles et les variables de contrôle.
func (s *Structure) init() {

	// Indique que le jeu est en cours d'exécution.
	s.Running = true

	// Initialise une liste vide pour stocker les lettres proposées par le joueur.
	s.Letter = []rune{}

	// Sélectionne un mot aléatoire pour le jeu et l'attribue à SecretWord.
	s.SecretWord = s.SelectRandomWord()

	// Initialise les tirets bas (ou espaces vides) représentant les lettres non encore devinées.
	s.Blanks = s.InitializeBlanks()

	// Définit le nombre de vies initiales du joueur.
	s.Lives = 10

	// Initialise une chaîne vide pour suivre les lettres déjà testées par le joueur.
	s.LetterTested = ""

	// Définit l'état de victoire à false (aucune victoire au début).
	s.Win = false

	// Définit l'état de défaite à false (aucune défaite au début).
	s.Lose = false
}
