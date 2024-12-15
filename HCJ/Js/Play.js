// Fonction pour ajouter une lettre au champ de saisie
function addLetter(letter) {
    const inputWord = document.getElementById('inputWord');
    inputWord.value += letter;
    inputWord.focus(); // Garde le focus sur le champ
}
    
// Met le focus automatique sur le champ texte au chargement de la page
window.onload = function () {
    const inputWord = document.getElementById('inputWord');
    inputWord.focus();
};
       
// Fonction pour ajouter une lettre au champ de saisie
function addLetter(letter) {
    const inputWord = document.getElementById('inputWord');
    inputWord.value += letter;
} ;
//LE PENDU
let essaies = {{.Essaies}}  // La valeur de .Essaies sera injectée ici depuis le serveur
const canvas = document.querySelector("canvas");
const ctx = canvas.getContext("2d");
canvas.width = 400;
canvas.height = 550;

// ** Configuration des lignes Retrowave **
ctx.beginPath();
ctx.lineCap = "round";
ctx.lineJoin = "round";
ctx.lineWidth = 8; // Épaisseur du contour

// Ajout d'effets lumineux pour les lignes
ctx.shadowColor = "rgba(255, 0, 255, 0.1)"; // Ombre néon
ctx.shadowBlur = 15; // Flou pour effet lumineux

// Couleur néon pour les lignes
ctx.strokeStyle = "#abfcff"; // Rose néon (peut être changé selon les goûts)

// ** Dessin des étapes du pendu 
if (essaies==9) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);

    ctx.stroke();

}else if (essaies==8) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);

    ctx.stroke();

}else if (essaies==7) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);

    ctx.stroke();

}else if (essaies==6) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);
    // Étape 4
    ctx.moveTo(280, 40);
    ctx.lineTo(280, 100);

    ctx.stroke();

}else if (essaies==5) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);
    // Étape 4
    ctx.moveTo(280, 40);
    ctx.lineTo(280, 100);
    // Étape 5
    ctx.moveTo(280, 140); 
    // Positionnement du cercle
    ctx.arc(280, 140, 40, 0, Math.PI * 2);

    ctx.stroke();

}else if (essaies==4) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);
    // Étape 4
    ctx.moveTo(280, 40);
    ctx.lineTo(280, 100);
    // Étape 5
    ctx.moveTo(280, 140); 
    // Positionnement du cercle
    ctx.arc(280, 140, 40, 0, Math.PI * 2);
    // Étape 6
    ctx.moveTo(280,180);
    ctx.lineTo(280, 400);

    ctx.stroke();

}else if (essaies==3) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);
    // Étape 4
    ctx.moveTo(280, 40);
    ctx.lineTo(280, 100);
    // Étape 5
    ctx.moveTo(280, 140); 
    // Positionnement du cercle
    ctx.arc(280, 140, 40, 0, Math.PI * 2);
    // Étape 6
    ctx.moveTo(280,180);
    ctx.lineTo(280, 400);
    // Étape 7
    ctx.moveTo(280, 400);
    ctx.lineTo(200, 460);

    ctx.stroke();

}else if (essaies==2) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);
    // Étape 4
    ctx.moveTo(280, 40);
    ctx.lineTo(280, 100);
    // Étape 5
    ctx.moveTo(280, 140); 
    // Positionnement du cercle
    ctx.arc(280, 140, 40, 0, Math.PI * 2);
    // Étape 6
    ctx.moveTo(280,180);
    ctx.lineTo(280, 400);
    // Étape 7
    ctx.moveTo(280, 400);
    ctx.lineTo(200, 460);
    // Étape 8
    ctx.moveTo(280, 400);
    ctx.lineTo(360, 460);
    
    ctx.stroke();

}else if (essaies==1) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);
    // Étape 4
    ctx.moveTo(280, 40);
    ctx.lineTo(280, 100);
    // Étape 5
    ctx.moveTo(280, 140); 
    // Positionnement du cercle
    ctx.arc(280, 140, 40, 0, Math.PI * 2);
    // Étape 6
    ctx.moveTo(280,180);
    ctx.lineTo(280, 400);
    // Étape 7
    ctx.moveTo(280, 400);
    ctx.lineTo(200, 460);
    // Étape 8
    ctx.moveTo(280, 400);
    ctx.lineTo(360, 460);
    // Étape 9
    ctx.moveTo(280, 250);
    ctx.lineTo(200, 310);

    ctx.stroke();

}else if (essaies==0) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // Étape 1
    ctx.moveTo(10, 530);
    ctx.lineTo(390, 530);
    // Étape 2
    ctx.moveTo(30, 530);
    ctx.lineTo(30, 40);
    // Étape 3
    ctx.moveTo(30, 40);
    ctx.lineTo(280, 40);
    ctx.moveTo(30, 90);
    ctx.lineTo(120, 40);
    // Étape 4
    ctx.moveTo(280, 40);
    ctx.lineTo(280, 100);
    // Étape 5
    ctx.moveTo(280, 140); 
    // Positionnement du cercle
    ctx.arc(280, 140, 40, 0, Math.PI * 2);
    // Étape 6
    ctx.moveTo(280,180);
    ctx.lineTo(280, 400);
    // Étape 7
    ctx.moveTo(280, 400);
    ctx.lineTo(200, 460);
    // Étape 8
    ctx.moveTo(280, 400);
    ctx.lineTo(360, 460);
    // Étape 9
    ctx.moveTo(280, 250);
    ctx.lineTo(200, 310);
    // Étape 10
    ctx.moveTo(280, 250);
    ctx.lineTo(360, 310);

    ctx.stroke();
}