package main

import (
    "fmt"
    "net/http"
    "hangmanweb"
)

func main() {
    // Configurer le gestionnaire pour la racine "/"
    http.HandleFunc("/", hangmanweb.Handler)
    fmt.Println("Serveur en écoute sur le port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Erreur lors du démarrage du serveur :", err)
    }
}
