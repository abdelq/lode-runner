const { connect } = require('net');
const tls = require('tls');
const fs = require('fs');

process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

// TODO Flag for port (process.argv)
//const client = tls.connect(443, {ca: [ fs.readFileSync('server.crt') ]}, () => {
//const client = tls.connect(1337, {}, () => {
const client = connect(1337, {}, () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);

    //var buffer = new Buffer(JSON.stringify({"name": name, "room": room}), "utf-8")
    let joinMsg = JSON.stringify({
        "event": "join",
        //"data": buffer
        "data": {"name": "guard", "room": "room", "role": 48}
    });
    console.log(joinMsg)
    client.write(joinMsg);
});

client.on('data', (data) => {
    let json = JSON.parse(data.toString());
    if (json.Event == "start")
        console.log(json.Data);
    else
        console.log(json.Event);
});
