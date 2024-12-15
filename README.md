
# Hangman web - Jeu du Pendu

Bienvenue dans **Hangman**, une version interactive du classique jeu du Pendu. Ce projet est développé avec **HTML**, **CSS**, **JavaScript** et **Golang**.

---

## 🎮 Fonctionnalités

- **Interface Utilisateur Interactive** : 
  - Utilisation d’un clavier virtuel pour entrer des lettres.
  - Utilisation du clavier physique pour entrer des lettres.
  - Mise à jour du mot à deviner, des essais restants et des lettres déjà utilisées.
- **Affichage du Pendu** :
  - Dessin progressif du pendu en fonction des erreurs.
  - Utilisation de la balise `<canvas>` pour le rendu graphique.
- **Gestion de Parties** :
  - Détection automatique de victoire ou défaite.
  - Redirection vers des pages spécifiques en cas de victoire ou défaite (`/Win` ou `/Lose`).
- **Prise en charge des données serveur** :
  - Dynamisme grâce à l'injection de variables comme `essaies`, `Motcachee`, et `LettresUtilisee`.
  
---

## 🚀 Installation

1. **Clonez le dépôt** :
   ```bash
   git clone https://github.com/<VotreNomUtilisateur>/hangman.git


**Utilisation**
- Démarrez le jeu depuis la page principale.
- Soumettez des lettres via :
    Le clavier virtuel.
    Le champ de texte prévu à cet effet.
- Observez vos progrès à travers l'affichage :
    Mot caché avec les lettres trouvées.
    Nombre d'essais restants.
    Dessin du pendu.

**Technologies Utilisées**
Frontend : HTML5, CSS3, JavaScript (ES6+).
Canvas API : Utilisée pour dessiner le pendu.
Backend Dynamique :
Variables serveur injectées dans le HTML.