var socket = new WebSocket("ws://" + location.host + "/ws");
var invalid = {};
var killed = {};
var maxLevel = {};

var startTime = window.performance.now();

socket.onopen = function () {
    socket.send(JSON.stringify({
        event: "list"
    }));
}

socket.onmessage = function (msg) {
    var messages = parseJSON(msg.data);

    if(!Array.isArray(messages))
        messages = [messages];

    messages.forEach(function(msg) {
        switch (msg.event) {
            case "start":
            case "next":
                if (rooms[msg.data.room] === undefined ||
                    invalid[msg.data.room] === undefined ||
                    invalid[msg.data.room] === true || msg.event == "start") {
                        draw(msg.data.tiles, msg.data.room, msg.data.lives, msg.data.level);
                } else {
                    redraw(msg.data.tiles, msg.data.room, msg.data.lives, msg.data.level);
                }
                invalid[msg.data.room] = false;
                rooms[msg.data.room] = msg.data.tiles;

                maxLevel[msg.data.room] = maxLevel[msg.data.room] || [msg.data.level, window.performance.now() - startTime];

                if(msg.data.level > maxLevel[msg.data.room][0]) {
                    console.log("aaa", msg.data);
                    maxLevel[msg.data.room] = [msg.data.level, window.performance.now() - startTime]
                }

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
                killed[msg.data] = window.performance.now() - startTime;

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
    });
}

function updateGrid() {
    var deads = document.querySelectorAll('canvas.dead').length;

    if(deads > 0)
        return;

    var len = document.querySelectorAll('canvas:not(.dead)').length;

    if(len == 0) {
        leaderboard();
        return;
    }

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

function leaderboard() {
    if(Object.keys(killed).length == 0)
        return;

    var block = document.createElement('div');

    var ranking = [];
    for(var team in maxLevel) {
        ranking.push([team, maxLevel[team]]);
    }

    ranking.sort(function(a, b) {
        if(a[1] < b[1])
            return 1;
        else if(a[1] > b[1])
            return -1;
        else {
            var time_a = a[1];
            var time_b = b[1];

            return time_a < time_b ? 1 : -1;
        }
    });

    block.innerHTML += '<h1>Leaderboard</h1><ul>' + ranking.map(function(x, i) {

        var symbol = ' ';

        if (i < 3)
            symbol = '★';

        return '<li class="leaderboard-' + i + '"><span class="symbol">' + symbol + '</span>' +
               ' #' + (i + 1) + ' - ' + x[0] +
               '<br/><small>       lvl: ' + maxLevel[x[0]][0] + ' - ' + Math.round(maxLevel[x[0]][1] * 100)/100000 + 's</small></li>';

    }).join('') + '</ul>';

    document.body.appendChild(block);
}
