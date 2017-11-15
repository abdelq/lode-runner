var directions = [null, "up", "left", "down", "right"];

var last_direction = "left";

function onkeypress(key) {
    var dir = directions.indexOf(key.name);

    if(dir > 0)
        last_direction = dir;

    if(key.name == "q")
        process.exit();
}

function start(data) {
    console.log(data);
}

function next(data) {
    // Random :
    /* var dir = Math.floor(4 * Math.random()) + 1;
       return {event: "move", direction: dir}; */

    // Default : dernière direction appuyée
    var dir = last_direction;
    last_direction = 0;
    return {event: "move", direction: dir};
}

// XXX Important : ne pas modifier ces lignes
module.exports.onkeypress = onkeypress;
module.exports.start = start;
module.exports.next = next;
