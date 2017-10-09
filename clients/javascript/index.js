const { connect } = require('net');
const { name, room } = require('./ai');

// TODO Flag for port (process.argv)
const client = connect(443, () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);
    // TODO Join a room
});
