/* Canvas */
var canvas = document.getElementById('canvas'),
    context = canvas.getContext('2d');

canvas.width = canvas.offsetWidth;
canvas.height = canvas.offsetHeight;

//context.imageSmoothingEnabled = false

/* Socket */
var socket = new WebSocket("ws://" + location.host + "/ws");

socket.onopen = function () {
    var room = location.hash.substr(1);

    socket.send(JSON.stringify({
        event: "join",
        data: {
            name: "spectator",
            room: room,
            role: 42
        }
    }));

    document.title = room + " - " + document.title
}

socket.onmessage = function (message) {
    var msg = JSON.parse(message.data);
    if (msg.event == "start")
        start(msg.data);
    else if (msg.event == "next")
        next(msg.data);
    else
        console.log(msg);
}

/* Game */
var tiles, tileHeight, tileWidth;

function start(data) {
    tiles = data.split("\n");

    tileHeight = canvas.height / tiles.length,
    tileWidth = canvas.width / tiles[0].length;

    context.clearRect(0, 0, canvas.width, canvas.height);
    for (var i = 0; i < tiles.length; i++) {
        for (var j = 0; j < tiles[i].length; j++) {
            context.drawImage(
                document.getElementById(tiles[i][j]),
                j * tileWidth, i * tileHeight,
                tileWidth, tileHeight
            );
        }
    }
}

function next(data) {
    if (tiles == null) {
        start(data);
        return;
    }

    var oldTiles = tiles;
    tiles = data.split("\n");

    for (var i = 0; i < tiles.length; i++) {
        for (var j = 0; j < tiles[i].length; j++) {
            if (tiles[i][j] != oldTiles[i][j]) {
                context.clearRect(
                    j * tileWidth, i * tileHeight,
                    tileWidth, tileHeight
                );
                context.drawImage(
                    document.getElementById(tiles[i][j]),
                    j * tileWidth, i * tileHeight,
                    tileWidth, tileHeight
                );
            }
        }
    }
}
