var keypress = require('keypress');

var { connect } = require('net');
var tls = require('tls');
var fs = require('fs');

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
    if (json.Event == "start")
        console.log(json.Data);
    else
        console.log(json.Event);
});

var directions = [
    null,
    "up",
    "left",
    "down",
    "right"
];

var last_direction = true;

process.stdin.on('keypress', function (ch, key) {
    /* console.log('got "keypress"', key); */

    var dir = directions.indexOf(key.name);

    if(dir > 0)
        send("move", {direction: dir});

    if([2,4].indexOf(dir) !== -1)
        last_direction = dir;

    if(key.name == "space")
        send("dig", {direction: last_direction});

    if(key.name == "q")
        process.exit();
});

process.stdin.setRawMode(true);
process.stdin.resume();
