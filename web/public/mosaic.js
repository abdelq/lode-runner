var socket = new WebSocket("ws://" + location.host + "/ws");
var invalid = {};

socket.onopen = function () {
    socket.send(JSON.stringify({
        event: "list"
    }));
}

socket.onmessage = function (msg) {
    msg = parseJSON(msg.data);
    switch (msg.event) {
        case "start":
        case "next":
            if (rooms[msg.data.room] === undefined ||
                invalid[msg.data.room] === undefined ||
                invalid[msg.data.room] === true || msg.event == "start") {
                    draw(msg.data.tiles, msg.data.room, msg.data.lives);
            } else {
                redraw(msg.data.tiles, msg.data.room, msg.data.lives);
            }
            invalid[msg.data.room] = false;
            rooms[msg.data.room] = msg.data.tiles;
            break;
        case "list":
            msg.data.forEach(function (room) {
                createCanvas(room, true);

                socket.send(JSON.stringify({
                    event: "join", data: { room: room, role: 42 }
                }));
            });
            updateGrid();
            break;
        case "quit":
            var canvas = document.getElementById(msg.data);
            var title = canvas.parentElement.querySelector('p');
            title.innerHTML = title.innerHTML.replace(/\(.\)/, '(dead)');
            title.style.color = 'gray';

            canvas.style = "filter: grayscale(100%)";
            canvas.classList.add('dead');
            setTimeout(function() {
                canvas.parentElement.remove();
                delete invalid[msg.data];
                updateGrid();
            }, 5000);
            break;
        default:
            console.log(msg.event + ": " + msg.data);
    }
}

function updateGrid() {
    var deads = document.querySelectorAll('canvas.dead').length;

    if(deads > 0)
        return;

    var len = document.querySelectorAll('canvas:not(.dead)').length;

    var cols;
    var rows;

    if(len > 16) {
        cols = rows = 5;
    } else if(len > 12) {
        cols = rows = 4;
    } else if(len > 9) {
        cols = 4;
        rows = 3;
    } else if(len > 6) {
        cols = rows = 3;
    } else if(len >= 5) {
        cols = 3;
        rows = 2;
    } else if(len >= 2) {
        cols = rows = 2;
    } else {
        cols = rows = 1;
    }

    var width = 100 / cols;
    var height = 100 / rows;

    document.body.innerHTML += '<style>.mosaic { width: ' + width + '%; height: ' + height + '%;}';

    for(var room in invalid) {
        invalid[room] = true;
    }
}
