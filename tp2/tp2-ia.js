// TODO : modifier cette valeur avec la clé secrète
var room = "";

/**
 * La fonction start() est appelée au début
 * d'un niveau et reçoit en paramètre la grille
 * initiale
 *
 * Les symboles sont :
 *   # : brique
 *   & : position initiale du joueur
 *   $ : sac d'or
 *   H : échelle
 *   - : corde
 *   S : sortie
 *   espace vide : rien de spécial sur cette case
 */
function start(data) {
    console.log(data);
}

/**
 * La fonction `next` est appelée automatiquement à
 * chaque tour et doit retourner un enregistrement
 * de la forme :
 *   {event: ..., direction: ...}
 *
 * Où : - event est un des deux événements "move", pour
 *        se déplacer ou "dig" pour creuser
 *      - direction est une des 4 directions haut/gauche/bas/droite,
 *        représenté par un nombre : 1 pour haut, 2 pour gauche, ...
 */
function next(data) {
    // TODO : Modifier ici
    // Envoyer une direction au hasard
    var dir = Math.floor(Math.random() * 4) + 1;
    return {event: "move", direction: dir};
}

// XXX Important : ne pas modifier ces lignes
module.exports.room = room;
module.exports.start = start;
module.exports.next = next;
