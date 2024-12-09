package hangman

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataForm struct {
	Motsecret       string
	Motcachee       string
	LettresUtilisee string
	Victoire        bool
	Essaies         int
	Echec           bool
}

func (s *Structure) web() {
	// chargement de tous les répertoirs présents dans "Hangman-Web"
	http.Handle("/hangmanstage/", http.StripPrefix("/hangmanstage/", http.FileServer(http.Dir("hangmanstage"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))
	http.Handle("/texte/", http.StripPrefix("/texte/", http.FileServer(http.Dir("texte"))))
	http.Handle("/HtmlCss/", http.StripPrefix("/HtmlCss/", http.FileServer(http.Dir("HtmlCss"))))

	http.HandleFunc("/", s.home)
	http.HandleFunc("/Home", s.home)
	http.HandleFunc("/Play", s.play)
	http.HandleFunc("/Lose", s.lose)
	http.HandleFunc("/Win", s.win)

	// chargement du port utilisé
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}

// fonctions pour chaque page

func (s *Structure) home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("HtmlCss/home.html"))
	tmpl.Execute(w, nil)
}

func (s *Structure) lose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HtmlCss/Lose.html"))
	tmpl.Execute(w, nil)
}

func (s *Structure) win(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HtmlCss/Win.html"))
	tmpl.Execute(w, nil)
}
func (s *Structure) play(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		letter := r.PostFormValue("letter")
		s.Letter = []rune(letter)
		var check bool = false

		if len(s.Letter) == 1 {
			if s.CheckLetter(s.Letter) {
				check = true
			}
		} else if len(s.Letter) == len(s.SecretWord) {
			check = s.CheckWord(s.Letter)
		} else if len(s.Letter) == 0 {
			check = true
		}

		if !check {
			s.Lives -= 1
		}

		s.CheckOut()

		web := DataForm{
			Motsecret:       s.ConvertRinS(s.SecretWord),
			Motcachee:       s.ConvertRinS(s.Blanks),
			LettresUtilisee: s.LetterTested,
			Victoire:        s.Win,
			Essaies:         s.Lives,
			Echec:           s.Lose,
		}

		tmpl := template.Must(template.ParseFiles("HtmlCss/Play.html"))
		tmpl.Execute(w, web)
	} else {
		web := DataForm{
			Motsecret:       s.ConvertRinS(s.SecretWord),
			Motcachee:       s.ConvertRinS(s.Blanks),
			LettresUtilisee: s.LetterTested,
			Victoire:        s.Win,
			Essaies:         s.Lives,
			Echec:           s.Lose,
		}

		tmpl := template.Must(template.ParseFiles("HtmlCss/Play.html"))
		tmpl.Execute(w, web)
	}
}
