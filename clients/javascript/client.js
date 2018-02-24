const net = require('net');
const JSONStream = require('JSONStream');
const Runner = require('./runner');

const runner = new Runner();
const client = net.connect(1337, "localhost");

client.on('connect', () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);
    client.write(JSON.stringify({
        event: "join",
        data: {
            name: runner.name,
            room: runner.room,
            level: runner.level
        }
    }));
});

client.pipe(JSONStream.parse()).on('data', (msg) => {
    if (msg.event === 'start')
        runner.start(msg.data);
    else if (msg.event === 'next')
        runner.next(msg.data.runner.x, msg.data.runner.y);
    else
        console.log(msg);
});

client.on('end', () => {
    console.log(`Connection to ${client.remoteAddress}:${client.remotePort} ended`);
});

client.on('error', (err) => {
    console.error(err.toString());
});

Runner.prototype.move = function(dir) {
    client.write(JSON.stringify({
        event: "move",
        data: {
            direction: dir,
            room: this.room
        }
    }));
}

Runner.prototype.dig = function(dir) {
    client.write(JSON.stringify({
        event: "dig",
        data: {
            direction: dir,
            room: this.room
        }
    }));
}
