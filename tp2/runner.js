if(process.argv.length == 2) {
    console.error('Usage: node ' + 'runner.js ' + '{clé secrète}');
    process.exit();
}

var room = process.argv[process.argv.length - 1];
var ip = "138.197.153.140";

const { connect } = require('net');

function send(event, data) {
    client.write(JSON.stringify({
        event: event,
        data: data
    }));
}

const client = connect(1337, ip, () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);
    send("join", {name: "runner", room: room});
});

process.stdin.setRawMode(true);
process.stdin.setEncoding("utf8");
process.stdin.on('data', (key) => {
    switch (key) {
        case 'w':
        case '\u001B\u005B\u0041':
            send("move", {direction: 1, room: room});
            break;
        case 'a':
        case '\u001B\u005B\u0044':
            send("move", {direction: 2, room: room});
            break;
        case 's':
        case '\u001B\u005B\u0042':
            send("move", {direction: 3, room: room});
            break;
        case 'd':
        case '\u001B\u005B\u0043':
            send("move", {direction: 4, room: room});
            break;
        case 'z':
            send("dig",  {direction: 2, room: room});
            break;
        case 'c':
            send("dig",  {direction: 4, room: room});
            break;
        case '\u0003':
            process.exit();
            break;
    }
});

