function reconnect() {
    var socket = new WebSocket("ws://" + location.host + "/ws");

    console.log(socket);

    socket.onopen = function () {
        var room = location.search.substr(1).split('=')[1];

        createCanvas(room);
        document.title = room + " - Lode Runner";

        socket.send(JSON.stringify({
            event: "join", data: { room: room, role: 42 }
        }));
    }

    socket.onmessage = function (msg) {
        var waiting = document.getElementById('waiting');
        if(waiting)
            waiting.style.display = 'none';

        var messages = parseJSON(msg.data);

        if(!Array.isArray(messages))
            messages = [messages];

        messages.forEach(function(msg) {


            switch (msg.event) {
                case "start":
                    var title = document.querySelector('p');
                    title.style.color = "";
                    title.innerHTML = "";
                case "next":
                    if (rooms[msg.data.room] === undefined || msg.event == "start") {
                        draw(msg.data.tiles, msg.data.room, msg.data.lives, msg.data.level);
                    } else {
                        redraw(msg.data.tiles, msg.data.room, msg.data.lives, msg.data.level);
                    }
                    rooms[msg.data.room] = msg.data.tiles;
                    break;
                case "quit":
                    var title = document.querySelector('p');
                    title.innerHTML = "Fin de la partie";
                    title.style.color = "#CC0000";
                    socket.close();
                    reconnect();
                    break;
                default:
                    console.log(msg.event + ": " + msg.data);
            }
        });
    }
}

reconnect();
