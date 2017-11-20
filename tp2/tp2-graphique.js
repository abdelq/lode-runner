function genererTable(width, height) {
    // TODO : remplir cette fonction pour générer un <table> HTML de la bonne taille
    return '<table id="table">' +
             '<tr>' +
               '<td id="0-0"></td><td id="0-1"></td><td id="0-2"></td><td id="0-3"></td><td id="0-4"></td>' +
             '</tr>' +
             '<tr>' +
               '<td id="1-0"></td><td id="1-1"></td><td id="1-2"></td><td id="1-3"></td><td id="1-4"></td>' +
             '</tr>' +
             '<tr>' +
               '<td id="2-0"></td><td id="2-1"></td><td id="2-2"></td><td id="2-3"></td><td id="2-4"></td>' +
             '</tr>' +
           '</table>';
}

function randomChoice(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}

function draw(map) {
    // Exemple de génération d'un tableau HTML avec des images aléatoires,
    // décommentez le code pour voir le résultat
    /*
    var container = document.getElementById('grid');

    var table = genererTable(0, 0);

    container.innerHTML = table;

    for(var i=0; i<3; i++) {
        for(var j=0; j<5; j++) {
            // NOTEZ : "#" a une signification spéciale dans les URLs.
            // Pour utiliser le path "img/#.png", on doit l'écrire sa forme encodée : img/%23.png
            var img = 'url("img/' + randomChoice(['%23', '-', '&', 'H', '$']) + '.png';

            document.getElementById(i + '-' + j).style.backgroundImage = img;
        }
    }
    */

    // Affichage temporaire de la grille en ASCII, pour vous aider à débugguer
    var gridAscii = document.getElementById('grid-ascii');
    gridAscii.innerHTML = map;
}
