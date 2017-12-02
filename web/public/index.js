/* Canvas */
var canvas  = document.getElementById('canvas'),
    context = canvas.getContext('2d');

canvas.width  = canvas.offsetWidth;
canvas.height = canvas.offsetHeight;

function drawLevel(tiles) {
    context.clearRect(0, 0, canvas.width, canvas.height);

    var tileHeight = canvas.height / tiles.length; // XXX
    var tileWidth = canvas.width / tiles[0].length; // XXX

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

/* Socket */
var socket = new WebSocket("ws://" + location.host + "/ws");

socket.onopen = function () {
    socket.send(JSON.stringify({
        event: "join",
        data: {
            name: "spectator",
            room: location.hash.substr(1),
            role: 42
        }
    }));
}

socket.onmessage = function (message) {
    var msg = JSON.parse(message.data);
    switch (msg.event) {
        case "start":
            drawLevel(msg.data.split("\n"));
            break;
        case "next":
            drawLevel(msg.data.split("\n")); // XXX
            break;
        default:
            console.log(msg);
            break;
    }
}
