var directions = [null, "up", "left", "down", "right", "space"];

var last_direction = 0;
var look_direction = 0;

function onkeypress(key) {
    var dir = directions.indexOf(key.name);

    if(dir > 0) {
        last_direction = dir;

        if(dir < 5)
            look_direction = dir;
    }

    if(key.name == "q")
        process.exit();
}

function start(data) {
    console.log(data);
}

function next(data) {
    // Envoie la dernière direction appuyée
    var dir = last_direction;
    last_direction = 0;

    if(dir == 5) {
        return {event: "dig", direction: look_direction};
    }

    return {event: "move", direction: dir};
}

// XXX Important : ne pas modifier ces lignes
module.exports.onkeypress = onkeypress;
module.exports.start = start;
module.exports.next = next;
