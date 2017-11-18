var images = {};

function draw(map) {
    var canvas = document.getElementById('c');
    var ctx = canvas.getContext('2d');

    var grid = map.split('\n').map(function(x) {
        return x.split('');
    });

    var h = grid.length;
    var w = grid[0].length;

    canvas.width = w * 32;
    canvas.height = h * 32;
    canvas.style.display = 'inline';
    document.getElementById('loading').style.display = '';

    grid.forEach(function(row, y) {
        row.forEach(function(e, x) {
            if(e in images)
                ctx.drawImage(images[e], x * 32, y * 32);
        });
    });
}

var ws = null;
var reconnect = true;

function start() {

    var room = document.getElementById('room').value;
    document.getElementById('c').style.display = 'none';

    if(ws) {
        reconnect = false;
        ws.close();
    }

    ws = new WebSocket('ws://159.203.8.35:1338/' + room);

    window.location.hash = '#' + room;

    ws.onmessage = function(msg) {
        reconnect = true;
        draw(msg.data);
    };

    ws.onerror = function(msg) {
        alert('Impossible de se connecter au serveur');
    }

    ws.onclose = function() {
        document.getElementById('loading').style.display = 'block';
        document.getElementById('c').style.display = '';
        if(reconnect)
            setTimeout(start, 500);
    };

    document.getElementById('loading').style.display = 'block';
    document.getElementById('joined-room').innerHTML = room;
}

document.addEventListener('DOMContentLoaded', function() {
    ["-","@","$","&","H","S","X","#"].forEach(function(e) {
        var img = new Image();
        img.src = "img/" + e.replace('#', '%23') + ".png";
        images[e] = img;
    });

    if(window.location.hash != '' && window.location.hash != '#') {
        var room = window.location.hash.slice(1);
        document.getElementById('room').value = room;
        start();
    }
});
