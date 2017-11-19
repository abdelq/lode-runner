/* Canvas */
var canvas = document.getElementById('canvas');
var context = canvas.getContext('2d');

canvas.width  = canvas.offsetWidth;
canvas.height = canvas.offsetHeight;

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
    if (msg.event != "start" && msg.event != "next") {
        console.log(msg);
        return;
    }

    context.clearRect(0, 0, canvas.width, canvas.height);

    var tiles = msg.data.split("\n");
    var tileHeight = canvas.height / tiles.length;
    var tileWidth = canvas.width / tiles[0].length;

    for (var i = 0; i < tiles.length; i++) {
        for (var j = 0; j < tiles[i].length; j++) {
            context.drawImage(
                document.getElementById(tiles[i][j]),
                j * tileWidth,
                i * tileHeight,
                tileWidth,
                tileHeight
            );
        }
    }
}
