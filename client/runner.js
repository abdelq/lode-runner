var room = "room";

const { connect } = require('net');

function send(event, data) {
    client.write(JSON.stringify({
        event: event,
        data: data
    }));
}

const client = connect(1337, "159.203.8.35", () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);
    send("join", {name: "runner", room: room});
});

client.on('data', (data) => {
    let msg = JSON.parse(data.toString())
    if (msg.event == "start" || msg.event == "next")
        console.log(msg.data);
    else
        console.log(msg)
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

