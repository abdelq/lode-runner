var {connect} = require('net');
var tls = require('tls');
var fs = require('fs');

var {room, start, next} = require('./tp2-ia.js');

if(room == "") {
    console.error('Veuillez entrer une clé secrète dans tp2-ia.js');
    process.exit();
}
    

var ip = "138.197.153.140";
process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

function send(event, data) {
    var msg = JSON.stringify({
        event: event,
        data: data
    });
    client.write(msg);
}

var client = connect("1337", ip, () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);
    send("join", {"name": "runner", "room": room});
});

client.on('data', (data) => {

    var json = {event: "none"};

    try {
        var json = JSON.parse(data.toString('utf8').trim());
    } catch(e) {
        console.error('JSON error : ', e, data.toString('utf8'));
    }

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
