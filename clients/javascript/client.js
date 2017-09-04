const WebSocket = require('ws');

const ws = new WebSocket('http://localhost:3000'); // TODO address as CLI argument

ws.on('open', function open() {
    // TODO
    ws.send(JSON.stringify({
        event: 'join',
        data: 'twado',
    }));
});

ws.on('message', function message(data) {
    console.log('message: ' + data);
});

ws.on('close', function close() {
    console.log('close');
});
