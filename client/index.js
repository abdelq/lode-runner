var keypress = require('keypress');

var {connect} = require('net');
var tls = require('tls');
var fs = require('fs');
var {onkeypress, start, next} = require('./tp1.js');
console.log(onkeypress, start, next);

keypress(process.stdin);

process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

function send(event, data) {
    var msg = JSON.stringify({
        event: event,
        data: data
    });
    console.log("SEND ", msg);
    client.write(msg);
}

// TODO Flag for port (process.argv)
//const client = tls.connect(443, {ca: [ fs.readFileSync('server.crt') ]}, () => {
//const client = tls.connect(1337, {}, () => {
var client = tls.connect(1337, {}, () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);
    send("join", {"name": "runner", "room": "room"});
});

client.on('data', (data) => {
    var json = JSON.parse(data.toString());

    var events = {
        start: start,
        next: next,
    };

    var out;

    if(json.Event in events) {
        out = events[json.Event](json.Data);
    }

    if(out !== undefined && "event" in out) {
        var event = out.event;
        delete out.event;
        send(event, out);
    }
});

process.stdin.on('keypress', function (ch, key) {
    onkeypress(key);
});

process.stdin.setRawMode(true);
process.stdin.resume();
