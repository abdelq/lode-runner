const WebSocket = require('ws');

const ws = new WebSocket('http://localhost:3000/ws');

ws.on('open', function open() {
    console.log('open');

    ws.send(JSON.stringify(
        {'type': 'join', 'data': 'twado'}
    ));
});

ws.on('message', function message(data) {
    console.log('message: ' + data);
});

ws.on('error', function error(data) {
    console.log('error: ' + data);
});

ws.on('close', function close() {
    console.log('close');
});
