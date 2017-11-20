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

function oneJSON(str) {
    var count = 0;

    var start = str.indexOf('{');
    if(start > -1)
        str = str.slice(start);

    for(var i=0; i<str.length; i++) {
        if (i !== 0 && count == 0)
            break;

        if(str[i] == '{') count++;
        else if(str[i] == '}') count--;
    }

    return str.slice(0, i);
}

client.on('data', (data) => {

    var json = {event: "none"};

    try {
        json = JSON.parse(data.toString('utf8').trim());
    } catch(e) {
        // XXX : Count Hackula strikes again !
        // Si on n'arrive pas à parser la string, on a possiblement
        // reçu plusieurs objets en une seule string... FIXME
        console.error('JSON error : ', e, data.toString('utf8'));

        json = JSON.parse(oneJSON(data.toString('utf8').trim()));
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
