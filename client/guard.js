const { connect } = require('net');

const client = connect(1337, {}, () => {
    console.log(`Connected to ${client.remoteAddress}:${client.remotePort}`);

    client.write(JSON.stringify({
        event: "join",
        data: {
            name: "guard",
            room: "room",
            role: "0".charCodeAt()
        }
    }));
});
