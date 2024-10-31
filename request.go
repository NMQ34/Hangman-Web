package main

import (
	"html/template"
	"net/http"
)

type Range struct {
	Letter string
}

type DataForm struct {
	Word           string
	BoolValue      bool
	NumberOfButton []Range
}

var Data DataForm

func main() {
	Alphabet()
	Data.Word = "Bonjour"
	Data.BoolValue = false
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/input", Input)
	http.ListenAndServe(":8080", nil)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, Data)
}

func Input(w http.ResponseWriter, r *http.Request) {
	Data.Word = r.FormValue("textearecup")
	Data.BoolValue = true

	MainPage(w, r)
}

func Alphabet() {
	for i := 'a'; i != 'z'+1; i++ {
		Data.NumberOfButton = append(Data.NumberOfButton, Range{Letter: string(i)})
	}
}
