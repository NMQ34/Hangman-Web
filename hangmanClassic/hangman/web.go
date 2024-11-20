package hangman

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataForm struct {
	Motsecret    string
	Motcachee    string
	Usingletters string
	Close        bool
	Win          bool
	Try          int
	End          bool
}

func (g *Structure) web() {
	// chargement de tous les répertoirs présents dans "Hangman-Web"
	http.Handle("/steady/", http.StripPrefix("/steady/", http.FileServer(http.Dir("steady"))))

	http.HandleFunc("/pageSpiderman", g.pageSpiderman)
	// chargement du port utilisé
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}

// fonctions pour chaque page

func (g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(g.pagesWeblist[0]))
	tmpl.Execute(w, nil)
}

func (g *Structure) pageHangman(w http.ResponseWriter, r *http.Request, indice int) {

	r.ParseForm() // permet de récupérer les caractères envoyés par les pages html
	tmpl := template.Must(template.ParseFiles(g.pagesWeblist[indice]))
	letter := r.Form.Get("letter") // initialisation de la lettre actuelle
	g.currentLetter = letter

	if len(g.currentLetter) > 0 { // si l'information reçue n'est pas vide( donc est une lettre)
		if g.currentLetter == "giveup" { // si l'information reçue est égale à "giveup", la partie est terminée car le joueur a abandonné
			g.end = true
		} else if g.verif(rune(g.currentLetter[0])) { // si la lettre n'a pas déja été utilisée : on l'ajoute aux lettres utilisées et on vérifie si elle est dans le mot
			g.usingLetters += "  " + string(g.currentLetter[0])
			g.letterTested = append(g.letterTested, string(g.currentLetter[0]))
			g.inWord(rune(g.currentLetter[0]))
		}
	}
	g.verifWin()     // on vérifie si le joueur a gagné la partie
	if g.try >= 10 { // si le joueur a eu + de 10 essaies, la partie est terminée
		g.end = true
	}
	if g.end { // si la partie est terminée et que l'utilisateur appuie sur le bouton "Menu" : On réinitialise tout et on relance la première page
		if g.currentLetter == "back" {
			g.closeWindow = true
		}
	}
	web := DrawWeb{
		Motsecret:    g.secretWord,
		Motcachee:    g.hiddenWord,
		Usingletters: g.usingLetters,
		Close:        g.closeWindow,
		Win:          g.win,
		Try:          g.try,
		End:          g.end,
	}
	tmpl.Execute(w, web) // permet d'envoyer les valeurs de la structure DrawWeb sur le html

	if g.closeWindow {
		g.init()
		return
	}
}
