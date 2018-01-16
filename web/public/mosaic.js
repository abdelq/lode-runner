var socket = new WebSocket("ws://" + location.host + "/ws");

socket.onopen = function () {
    socket.send(JSON.stringify({
        event: "list"
    }));
}

socket.onmessage = function (msg) {
    msg = JSON.parse(msg.data);
    switch (msg.event) {
        case "start":
        case "next":
            if (rooms[msg.data.room] === undefined) {
                draw(msg.data.tiles, msg.data.room);
            } else {
                redraw(msg.data.tiles, msg.data.room);
            }
            rooms[msg.data.room] = msg.data.tiles;
            break;
        case "list":
            msg.data.forEach(function (room) {
                createCanvas(room, true);

                socket.send(JSON.stringify({
                    event: "join", data: { room: room, role: 42 }
                }));
            });
            break;
        case "quit":
            var canvas = document.getElementById(msg.data);
            canvas.style = "filter: grayscale(100%)";
            setTimeout(function() { canvas.remove(); }, 5000);
            break;
        default:
            console.log(msg.event + ": " + msg.data);
    }
}
