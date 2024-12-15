package hangman

import (
	"fmt"
	"html/template"
	"net/http"
)

// DataForm contient les données utilisées pour lier la logique du jeu aux pages HTML.
type DataForm struct {
	// Motsecret : Le mot secret à deviner.
	Motsecret string
	// Motcachee : Le mot masqué visible par le joueur (ex: "_ _ _ _").
	Motcachee string
	// LettresUtilisee : Les lettres déjà testées par le joueur.
	LettresUtilisee string
	// Victoire : Indique si le joueur a gagné.
	Victoire bool
	// Essaies : Nombre de vies restantes pour le joueur.
	Essaies int
	// Echec : Indique si le joueur a perdu.
	Echec bool
}

// web configure le serveur HTTP et les routes nécessaires pour le jeu.
func (s *Structure) web() {
	// Chargement des fichiers statiques (images, texte, HTML/CSS/JS).
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))
	http.Handle("/texte/", http.StripPrefix("/texte/", http.FileServer(http.Dir("texte"))))
	http.Handle("/HCJ/", http.StripPrefix("/HCJ/", http.FileServer(http.Dir("HCJ"))))

	// Définition des routes pour les différentes pages du jeu.
	http.HandleFunc("/", s.Home)      // Page d'accueil
	http.HandleFunc("/Home", s.Home) // Alias pour la page d'accueil
	http.HandleFunc("/Play", s.play) // Page de jeu
	http.HandleFunc("/Lose", s.lose) // Page de défaite
	http.HandleFunc("/Win", s.win)   // Page de victoire
	http.HandleFunc("/Reset", s.reset) // Réinitialisation du jeu

	// Démarrage du serveur HTTP sur le port 8080.
	fmt.Println("http://localhost:8080/") // Indique où accéder au jeu.
	http.ListenAndServe(":8080", nil)
}

// Home gère la page d'accueil du jeu.
func (s *Structure) Home(w http.ResponseWriter, r *http.Request) {
	// Charge et affiche le fichier HTML correspondant à la page d'accueil.
	tmpl := template.Must(template.ParseFiles("./HCJ/Html/Home.html"))
	tmpl.Execute(w, nil)
}

// lose gère l'affichage de la page de défaite.
func (s *Structure) lose(w http.ResponseWriter, r *http.Request) {
	// Charge le fichier HTML pour la page de défaite et transmet les données nécessaires.
	tmpl := template.Must(template.ParseFiles("./HCJ/Html/Lose.html"))
	web := DataForm{
		Essaies:   s.Lives,                          // Nombre de vies restantes.
		Motsecret: s.ConvertRinS(s.SecretWord),      // Mot à deviner.
	}
	tmpl.Execute(w, web)
}

// win gère l'affichage de la page de victoire.
func (s *Structure) win(w http.ResponseWriter, r *http.Request) {
	// Charge le fichier HTML pour la page de victoire et transmet les données nécessaires.
	tmpl := template.Must(template.ParseFiles("./HCJ/Html/Win.html"))
	web := DataForm{
		Essaies:   s.Lives,                          // Nombre de vies restantes.
		Motcachee: s.ConvertRinS(s.Blanks),          // Mot deviné par le joueur.
	}
	tmpl.Execute(w, web)
}

// play gère la logique du jeu et l'affichage de la page principale.
func (s *Structure) play(w http.ResponseWriter, r *http.Request) {
	// Vérifie si la requête est une méthode POST (soumission de lettre ou mot).
	if r.Method == "POST" {
		r.ParseForm() // Parse les données envoyées via le formulaire.
		letter := r.PostFormValue("letter") // Récupère la lettre ou le mot envoyé.
		s.Letter = []rune(letter)          // Convertit la lettre/mot en tableau de runes.
		var check bool = false

		// Vérifie si la saisie correspond à une lettre, un mot ou est vide.
		if len(s.Letter) == 1 {
			// Vérifie si la lettre est correcte.
			if s.CheckLetter(s.Letter) {
				check = true
			}
		} else if len(s.Letter) == len(s.SecretWord) {
			// Vérifie si le mot complet est correct.
			check = s.CheckWord(s.Letter)
		} else if len(s.Letter) == 0 {
			// Si aucune saisie, considère comme valide pour éviter une pénalité.
			check = true
		}

		// Si la saisie est incorrecte, réduit le nombre de vies.
		if !check {
			s.Lives -= 1
		}

		// Vérifie si le joueur a gagné ou perdu.
		s.CheckOut()

		// Met à jour les données à transmettre à la page HTML.
		web := DataForm{
			Motsecret:       s.ConvertRinS(s.SecretWord), // Mot à deviner.
			Motcachee:       s.ConvertRinS(s.Blanks),    // Mot masqué visible par le joueur.
			LettresUtilisee: s.LetterTested,             // Lettres déjà testées.
			Victoire:        s.Win,                      // Indique si le joueur a gagné.
			Essaies:         s.Lives,                   // Nombre de vies restantes.
			Echec:           s.Lose,                    // Indique si le joueur a perdu.
		}

		// Charge et affiche la page de jeu.
		tmpl := template.Must(template.ParseFiles("HCJ/Html/Play.html"))
		tmpl.Execute(w, web)
	} else {
		// Si la méthode est GET, affiche la page sans modification.
		web := DataForm{
			Motsecret:       s.ConvertRinS(s.SecretWord), // Mot à deviner.
			Motcachee:       s.ConvertRinS(s.Blanks),    // Mot masqué visible par le joueur.
			LettresUtilisee: s.LetterTested,             // Lettres déjà testées.
			Victoire:        s.Win,                      // Indique si le joueur a gagné.
			Essaies:         s.Lives,                   // Nombre de vies restantes.
			Echec:           s.Lose,                    // Indique si le joueur a perdu.
		}

		// Charge et affiche la page de jeu.
		tmpl := template.Must(template.ParseFiles("HCJ/Html/Play.html"))
		tmpl.Execute(w, web)
	}
}

// reset gère la réinitialisation du jeu.
func (s *Structure) reset(w http.ResponseWriter, r *http.Request) {
	s.Reset() // Réinitialise toutes les variables du jeu.
	http.Redirect(w, r, "/Play", http.StatusSeeOther) // Redirige vers la page de jeu.
}
