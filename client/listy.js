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
});

setInterval(function() {
    send("listylist", {});
}, 1000);

client.on('data', (data) => {
    console.log(JSON.parse(data).data.toString());
});

