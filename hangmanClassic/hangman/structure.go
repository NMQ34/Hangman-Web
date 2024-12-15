package hangman


// Structure représente l'état et les données nécessaires au fonctionnement du jeu du pendu.
type Structure struct {
	// Running indique si le jeu est actuellement en cours.
	Running bool

	// Variables liées à la logique du jeu :
	// Letter : Les lettres actuellement proposées par le joueur.
	Letter []rune
	// SecretWord : Le mot à deviner, stocké sous forme de tableau de runes.
	SecretWord []rune
	// Blanks : Les emplacements visibles pour le joueur, initialisés en tirets bas et remplis au fur et à mesure.
	Blanks []rune
	// Lives : Le nombre de vies restantes pour le joueur.
	Lives int
	// LetterTested : Les lettres déjà testées, sous forme de chaîne de caractères (affichées au joueur).
	LetterTested string

	// Variables d'état qui indiquent si le joueur a gagné ou perdu :
	Win  bool // True si le joueur a gagné.
	Lose bool // True si le joueur a perdu.
}

// Run est la fonction principale pour exécuter le jeu.
// Elle initialise le jeu et lance l'interface Web si le jeu est en cours (`Running` est true).
func (s *Structure) Run() {
	// Initialise les variables du jeu (mot, vies, etc.).
	s.init()
	// Si le jeu est actif, lance l'interface Web.
	if s.Running {
		s.web()
	}
}

// Reset permet de réinitialiser le jeu à son état initial.
// Cette fonction appelle la méthode `init` pour réinitialiser toutes les variables.
func (s *Structure) Reset() {
	s.init()
}
