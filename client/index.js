var {connect} = require('net');
var tls = require('tls');
var fs = require('fs');

if (process.argv.indexOf('--clavier') > -1) {
    var keypress = require('keypress');

    var {onkeypress, start, next} = require('./clavier.js');

    keypress(process.stdin);
        process.stdin.on('keypress', function (ch, key) {
        onkeypress(key);
    });

    process.stdin.setRawMode(true);
    process.stdin.resume();
} else {
    var {start, next} = require('./tp2.js');
}

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
var client = connect("1337", "159.203.8.35", () => {
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

    if(json.event in events) {
        out = events[json.event](json.data);
    }

    if(out !== undefined && "event" in out) {
        var event = out.event;
        delete out.event;
        send(event, out);
    }
});
