var socket = new WebSocket("ws://localhost:8080/ws"); // XXX

socket.addEventListener('open', function (event) {
    socket.send(JSON.stringify({
        event: "join",
        data: {name: "player", room: "room"}
    }));
});

socket.addEventListener('message', function (event) {
    console.log(event.data);
});
