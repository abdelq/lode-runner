var socket = new WebSocket("ws://" + location.host + "/ws");

socket.onopen = function () {
    var room = new URLSearchParams(location.search).get("name");

    createCanvas(room);
    document.title = room + " - " + document.title;

    socket.send(JSON.stringify({
        event: "join", data: { room: room, role: 42 }
    }));
}

socket.onmessage = function (msg) {
    msg = JSON.parse(msg.data);
    switch (msg.event) {
        case "start":
        case "next":
            if (rooms[msg.data.room] === undefined) {
                draw(msg.data.tiles, msg.data.room, msg.data.lives);
            } else {
                redraw(msg.data.tiles, msg.data.room, msg.data.lives);
            }
            rooms[msg.data.room] = msg.data.tiles;
            break;
        case "quit":
            console.log(msg.data);
            socket.close();
            break;
        default:
            console.log(msg.event + ": " + msg.data);
    }
}
