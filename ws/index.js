var WebSocket = require('ws');
var {connect} = require('net');
var tls = require('tls');
var fs = require('fs');
var url = require('url');

process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

var wss = new WebSocket.Server({ port: 1338 });
var ip = "localhost";

wss.on('connection', function connection(ws, req) {
    var location = url.parse(req.url, true);

    // open tcp client
    var room = location.pathname.slice(1);

    if(!room.length) {
        ws.close();
        return;
    }

    var client = connect(1337, ip, () => {
	if(!client) return;

        console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);
        send(client, "join", {name: "spectator", room: room, role: 42});
    });

    client.on('data', (data) => {
        try {
            var grid = JSON.parse(data.toString()).data;
        } catch(e) {
            var grid = data.toString();
            console.log('JSON error parsing :');
            console.log(grid);
        }
        if (ws.readyState === WebSocket.OPEN) {
            ws.send(grid);
        }
    });

    client.on('close', function() {
        ws.close();
    });

    ws.on('close', function() {
        // TODO : Kill tcp client
        client = null;
    });
});

function send(client, event, data) {
    var msg = JSON.stringify({
        event: event,
        data: data
    });
    console.log('SEND', msg);
    client.write(msg);
}

