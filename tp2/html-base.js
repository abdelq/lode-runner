var ws = null;
var reconnect = true;
var ip = "138.197.153.140";

function start() {

    var room = document.getElementById('room').value;

    if(ws) {
        reconnect = false;
        ws.close();
    }

    ws = new WebSocket('ws://' + ip + ':1338/' + room);

    window.location.hash = '#' + room;

    ws.onmessage = function(msg) {
        reconnect = true;
        draw(msg.data);
        document.getElementById('loading').style.display = '';
    };

    ws.onerror = function(msg) {
        alert('Impossible de se connecter au serveur');
    };

    ws.onclose = function() {
        document.getElementById('loading').style.display = 'block';
        if(reconnect)
            setTimeout(start, 500);
    };

    document.getElementById('loading').style.display = 'block';
    document.getElementById('joined-room').innerHTML = room;
}

document.addEventListener('DOMContentLoaded', function() {
    if(window.location.hash != '' && window.location.hash != '#') {
        var room = window.location.hash.slice(1);
        document.getElementById('room').value = room;
        start();
    }
});
