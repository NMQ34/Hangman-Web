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
	http.HandleFunc("/home", s.home)
	http.HandleFunc("/input", s.game)
	http.HandleFunc("/lose", s.lose)
	http.HandleFunc("/win", s.win)

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

}

func (s *Structure) win(w http.ResponseWriter, r *http.Request) {

}
func (s *Structure) game(w http.ResponseWriter, r *http.Request) {
	fmt.Println("debut")

	if r.Method == "POST" {
		r.ParseForm()
		letter := r.PostFormValue("letter")
		fmt.Println(letter, " --------------> la letter ou mot qu'on recoit")

		s.Letter = []rune(letter)
		fmt.Println(s.Letter, " --------------> la letter ou mot qu'on recoit en rune")

		var check bool = false
		fmt.Println("avant")

		if len(s.Letter) == 1 {
			fmt.Println("first")

			if s.CheckLetter(s.Letter) {
				fmt.Println("1")
				s.VerifLetter(s.Letter)
				check = true
			}

		} else if len(s.Letter) == len(s.SecretWord) {
			fmt.Println("2")
			s.CheckWord(s.Letter)
			check = true

		} else if len(s.Letter) == 0 {
			fmt.Println("3")
			check = true
		}
		fmt.Println("derbierr")
		if !check {
			fmt.Println("perdu")
			s.Lives -= 1
		}
		fmt.Println("a checkout")
		s.CheckOut()
		fmt.Println("ap checkout")
		web := DataForm{
			Motsecret:       s.ConvertRinS(s.SecretWord),
			Motcachee:       s.ConvertRinS(s.Blanks),
			LettresUtilisee: s.LetterTested,
			Victoire:        s.Win,
			Essaies:         s.Lives,
			Echec:           s.Lose,
		}
		fmt.Println("----------------deuxieme atemp")
		fmt.Println(web.Motsecret)
		fmt.Println(web.Motcachee)
		fmt.Println(web.LettresUtilisee)
		fmt.Println(web.Victoire)
		fmt.Println(web.Essaies)
		fmt.Println(web.Echec)

		tmpl := template.Must(template.ParseFiles("HtmlCss/game.html"))
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
		fmt.Println("--------------premier attemp")
		fmt.Println(web.Motsecret)
		fmt.Println(web.Motcachee)
		fmt.Println(web.LettresUtilisee)
		fmt.Println(web.Victoire)
		fmt.Println(web.Essaies)
		fmt.Println(web.Echec)

		tmpl := template.Must(template.ParseFiles("HtmlCss/game.html"))
		tmpl.Execute(w, web)
	}
}
